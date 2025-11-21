// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Test, console} from "forge-std/Test.sol";
import {IERC20Errors} from "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";
import {Ownable} from "openzeppelin-contracts/contracts/access/Ownable.sol";
import {CrosschainNubilaNetwork} from "../src/CrosschainNubilaNetwork.sol";

contract CrosschainNubilaNetworkTest is Test {
    CrosschainNubilaNetwork public token;
    address public owner;
    address public minter1;
    address public minter2;
    address public user1;
    address public user2;

    event MinterAdded(address indexed account);
    event MinterRemoved(address indexed account);
    event CrosschainMinted(address indexed receiver, uint256 amount);
    event CrosschainBurned(address indexed sender, address indexed receiver, uint256 amount);

    function setUp() public {
        owner = makeAddr("owner");
        minter1 = makeAddr("minter1");
        minter2 = makeAddr("minter2");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        vm.startPrank(owner);
        token = new CrosschainNubilaNetwork();
        vm.stopPrank();
    }

    function testInitialState() public view {
        assertEq(token.name(), "Nubila Network");
        assertEq(token.symbol(), "NB");
        assertEq(token.totalSupply(), 0);
        assertEq(token.owner(), owner);
        assertFalse(token.isMinter(owner));
        assertFalse(token.isMinter(minter1));
    }

    function testAddMinter() public {
        vm.expectEmit(true, true, true, true);
        emit MinterAdded(minter1);
        vm.startPrank(owner);
        token.addMinter(minter1);
        assertTrue(token.isMinter(minter1));
        vm.expectRevert(abi.encodeWithSelector(CrosschainNubilaNetwork.ErrAlreadyMinter.selector, minter1));
        token.addMinter(minter1); // Should revert if already a minter
        vm.expectRevert(CrosschainNubilaNetwork.ErrZeroAddress.selector);
        token.addMinter(address(0)); // Should revert for zero address
        vm.stopPrank();

        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, user1));
        vm.prank(user1);
        token.addMinter(minter2); // Only owner can add minters
    }

    function testRemoveMinter() public {
        vm.startPrank(owner);
        token.addMinter(minter1);
        assertTrue(token.isMinter(minter1));

        token.removeMinter(minter1);
        assertFalse(token.isMinter(minter1));
        vm.expectRevert(abi.encodeWithSelector(CrosschainNubilaNetwork.ErrNotMinter.selector, minter1));
        token.removeMinter(minter1); // Should revert if not a minter
        vm.stopPrank();

        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, user1));
        vm.prank(user1);
        token.removeMinter(minter2); // Only owner can remove minters
    }

    function testMint() public {
        vm.startPrank(owner);
        token.addMinter(minter1);
        vm.stopPrank();

        vm.expectEmit(true, true, true, true);
        emit CrosschainMinted(user1, 100 ether);
        vm.prank(minter1);
        token.crosschainMint(user1, 100 ether);
        assertEq(token.balanceOf(user1), 100 ether);
        assertEq(token.totalSupply(), 100 ether);

        vm.expectRevert(abi.encodeWithSelector(CrosschainNubilaNetwork.ErrNotAuthorizedMinter.selector, user2));
        vm.prank(user2);
        token.crosschainMint(user2, 50 ether); // Only minters can mint
    }

    function testCrosschainBurn() public {
        vm.startPrank(owner);
        token.addMinter(minter1);
        vm.stopPrank();

        vm.expectEmit(true, true, true, true);
        emit CrosschainMinted(user1, 200 ether);
        vm.prank(minter1);
        token.crosschainMint(user1, 200 ether);
        assertEq(token.balanceOf(user1), 200 ether);

        vm.startPrank(user1);
        vm.expectEmit(true, true, true, true);
        emit CrosschainBurned(user1, user2, 50 ether);
        token.crosschainBurn(50 ether, user2);
        assertEq(token.balanceOf(user1), 150 ether);
        assertEq(token.balanceOf(user2), 0); // Bridged tokens are burned, not transferred to receiver
        assertEq(token.totalSupply(), 150 ether);
        vm.expectEmit(true, true, true, true);
        emit CrosschainBurned(user1, user1, 25 ether);
        token.crosschainBurn(25 ether, address(0)); // Test with zero address receiver, should burn from sender
        assertEq(token.balanceOf(user1), 125 ether); // This line is duplicated, but it's fine for the test
        assertEq(token.totalSupply(), 125 ether);

        vm.expectRevert(abi.encodeWithSelector(IERC20Errors.ERC20InsufficientBalance.selector, user1, 125 ether, 200 ether));
        token.crosschainBurn(200 ether, user2); // Should revert if amount exceeds balance
        vm.stopPrank();
    }
}