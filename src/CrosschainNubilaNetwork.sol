// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";
import "openzeppelin-contracts/contracts/access/Ownable.sol";

contract CrosschainNubilaNetwork is ERC20, Ownable {
    modifier onlyMinter {
        if (!minters[msg.sender]) {
            revert ErrNotAuthorizedMinter(msg.sender);
        }
        _;
    }
    error ErrZeroAddress();
    error ErrAlreadyMinter(address account);
    error ErrNotAuthorizedMinter(address account);
    error ErrNotMinter(address account);

    event MinterAdded(address indexed account);
    event MinterRemoved(address indexed account);
    event CrosschainMinted(address indexed receiver, uint256 amount);
    event CrosschainBurned(address indexed sender, address indexed receiver, uint256 amount);

    mapping (address => bool) private minters;

    constructor() ERC20("Nubila Network", "NB") Ownable(msg.sender) {
    }

    function addMinter(address account) public onlyOwner {
        if (account == address(0)) {
            revert ErrZeroAddress();
        }
        if (minters[account]) {
            revert ErrAlreadyMinter(account);
        }
        minters[account] = true;
        emit MinterAdded(account);
    }

    function removeMinter(address account) public onlyOwner {
        if (!minters[account]) {
            revert ErrNotMinter(account);
        }
        minters[account] = false;
        emit MinterRemoved(account);
    }

    function isMinter(address account) public view returns (bool) {
        return minters[account];
    }

    function crosschainMint(address to, uint256 amount) public onlyMinter {
        _mint(to, amount);
        emit CrosschainMinted(to, amount);
    }

    function crosschainBurn(uint256 amount, address receiver) public {
        address sender = msg.sender;
        if (receiver == address(0)) {
            receiver = sender;
        }
        _burn(sender, amount);
        emit CrosschainBurned(sender, receiver, amount);
    }
}
