pragma solidity ^0.8.0;

import {IAny2EVMMessageReceiver} from "../../interfaces/IAny2EVMMessageReceiver.sol";

import {EVM2EVMOnRampSetup} from "../onRamp/EVM2EVMOnRampSetup.t.sol";
import {CCIPClientExample} from "../../applications/CCIPClientExample.sol";
import {Client} from "../../libraries/Client.sol";

import {ERC165Checker} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/introspection/ERC165Checker.sol";
import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract CCIPClientExample_sanity is EVM2EVMOnRampSetup {
  function testExamples() public {
    CCIPClientExample exampleContract = new CCIPClientExample(s_sourceRouter, IERC20(s_sourceFeeToken));
    deal(address(exampleContract), 100 ether);
    deal(s_sourceFeeToken, address(exampleContract), 100 ether);

    // feeToken approval works
    assertEq(IERC20(s_sourceFeeToken).allowance(address(exampleContract), address(s_sourceRouter)), 2 ** 256 - 1);

    // Can set chain
    Client.EVMExtraArgsV1 memory extraArgs = Client.EVMExtraArgsV1({gasLimit: 300_000});
    bytes memory encodedExtraArgs = Client._argsToBytes(extraArgs);
    exampleContract.enableChain(DEST_CHAIN_SELECTOR, encodedExtraArgs);
    assertEq(exampleContract.s_chains(DEST_CHAIN_SELECTOR), encodedExtraArgs);

    address toAddress = address(100);

    // Can send data pay native
    exampleContract.sendDataPayNative(DEST_CHAIN_SELECTOR, abi.encode(toAddress), bytes("hello"));

    // Can send data pay feeToken
    exampleContract.sendDataPayFeeToken(DEST_CHAIN_SELECTOR, abi.encode(toAddress), bytes("hello"));

    // Can send data tokens
    assertEq(
      address(s_onRamp.getPoolBySourceToken(DEST_CHAIN_SELECTOR, IERC20(s_sourceTokens[1]))),
      address(s_sourcePools[1])
    );
    deal(s_sourceTokens[1], OWNER, 100 ether);
    IERC20(s_sourceTokens[1]).approve(address(exampleContract), 1 ether);
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({token: s_sourceTokens[1], amount: 1 ether});
    exampleContract.sendDataAndTokens(DEST_CHAIN_SELECTOR, abi.encode(toAddress), bytes("hello"), tokenAmounts);
    // Tokens transferred from owner to router then burned in pool.
    assertEq(IERC20(s_sourceTokens[1]).balanceOf(OWNER), 99 ether);
    assertEq(IERC20(s_sourceTokens[1]).balanceOf(address(s_sourceRouter)), 0);

    // Can send just tokens
    IERC20(s_sourceTokens[1]).approve(address(exampleContract), 1 ether);
    exampleContract.sendTokens(DEST_CHAIN_SELECTOR, abi.encode(toAddress), tokenAmounts);

    // Can receive
    assertTrue(ERC165Checker.supportsInterface(address(exampleContract), type(IAny2EVMMessageReceiver).interfaceId));

    // Can disable chain
    exampleContract.disableChain(DEST_CHAIN_SELECTOR);
  }
}
