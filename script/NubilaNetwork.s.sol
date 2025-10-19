// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {NubilaNetwork} from "../src/NubilaNetwork.sol";
import {VestingManager} from "../src/VestingManager.sol";
contract NubilaNetworkScript is Script {

    function setUp() public {}

    function run() public {
        vm.startBroadcast();
        uint256 tge = uint256(vm.envUint("TGE"));
        console.log("TGE read from env:", tge);
        NubilaNetwork nubilaNetwork = new NubilaNetwork(msg.sender);
        console.log("NubilaNetwork deployed to:", address(nubilaNetwork));
        address[] memory beneficiaries = new address[](10);
        beneficiaries[0] = address(0xB2C2afD942230a24AE2ffd33026F7dc925d2F384); // Device
        beneficiaries[1] = address(0x5c846C7ED2199511c14ac7613fC7cab4862E06d1); // Node
        beneficiaries[2] = address(0x1EAd8d91573884d051fCB9dFe9309f586603B5Cb); // Backer1
        beneficiaries[3] = address(0xeDff823C45Bb1Ccf7eBD181444534163b9Cf1560); // Backer2
        beneficiaries[4] = address(0xa526BfE1C2C6392e3860E0C0B20F23a7c5c6b236); // POS
        beneficiaries[5] = address(0x90933E2047E223408cf76F581E4afF9615fFfb88); // Treasury
        beneficiaries[6] = address(0x37B09e62Db01C74CBEF8a57c95D41A8AF05e3926); // Team & Advisor
        beneficiaries[7] = address(0x22AEb8A587E367C8341C8d3ad1A324a1B644A1E1); // Liquidity
        beneficiaries[8] = address(0xc8FF467B4405676cD8E3a4bF6E108eDbB4aB36EE); // Airdrop
        beneficiaries[9] = address(0x247a5B9b4a74137130FD0165B006174a76494C68); // Marketing
        VestingManager vestingManager = new VestingManager(address(nubilaNetwork), tge, beneficiaries);
        console.log("VestingManager deployed to:", address(vestingManager));
        vm.assertTrue(nubilaNetwork.transfer(address(vestingManager), nubilaNetwork.balanceOf(msg.sender)));

        vm.stopBroadcast();
    }
}
