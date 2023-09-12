// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import {Ownable} from "../lib/openzeppelin-contracts/contracts/access/Ownable.sol";

/// @title  PriceOracle
/// @notice Report proof and submit result after offchain validation
contract OptimisticOracle is Ownable {
    /// -----------------------------------------------------------------------
    /// Error
    /// -----------------------------------------------------------------------

    error TDSContractIsDefault();
    error TDSContractUnderReporting();
    error TDSContractReportTimeout();
    error TDSContractAlreadyReported();
    error ReportNotFound();

    /// -----------------------------------------------------------------------
    /// Structure Storage
    /// -----------------------------------------------------------------------

    struct Report {
        uint256 tdsContractId;
        uint256 reportId;
        uint64 startTime;
        uint32 referenceEvent;
        bytes triggerType;
        bytes proof;
    }

    struct Result {
        uint256 tdsContractId;
        uint256 reportId;
        uint64 submittedAt;
        bool isDefault;
        bytes result;
    }

    /// -----------------------------------------------------------------------
    /// Events
    /// -----------------------------------------------------------------------

    event ReportCreated(
        uint256 indexed tdsContractId,
        uint256 indexed reportId,
        uint64 indexed startTime,
        uint32 referenceEvent,
        bytes triggerType,
        bytes proof
    );

    event ReportResultSubmitted(
        uint256 indexed tdsContractId,
        uint256 indexed reportId,
        uint64 indexed submittedAt,
        bool isDefault,
        bytes result
    );

    /// -----------------------------------------------------------------------
    /// Mutable Storage
    /// -----------------------------------------------------------------------

    // solhint-disable-next-line
    uint64 internal REPORTING_PERIOD = 3 * 86400; // count in second

    uint256 internal _lastReportId;

    mapping(uint256 => Report) internal _reportManager; // reportId => report

    mapping(uint256 => Result) internal _resultManager; // reportId => report

    mapping(uint256 => uint256[]) internal _tdsContractReport; // tdsContract => list of reportId

    mapping(uint256 => uint256) internal _tdsContractDefaultReportId; // tdsContractId => reportId of first submission result that confirm default

    constructor() Ownable(msg.sender){}
    /// -----------------------------------------------------------------------
    /// Modifier
    /// -----------------------------------------------------------------------

    modifier canReport(uint256 tdsContractId) {
        if (_tdsContractReport[tdsContractId].length > 0) {
            uint256 lastReport = _tdsContractReport[tdsContractId][
                _tdsContractReport[tdsContractId].length - 1
                ];
            if (_resultManager[lastReport].isDefault)
                revert TDSContractIsDefault();
            if (
                _reportManager[lastReport].startTime + REPORTING_PERIOD >
                block.timestamp
            ) revert TDSContractUnderReporting();
        }
        _;
    }

    modifier canSubmitResult(uint256 reportId) {
        if (_reportManager[reportId].tdsContractId == 0)
            revert ReportNotFound();
        if (
            _reportManager[reportId].startTime + REPORTING_PERIOD <
            block.timestamp
        ) revert TDSContractReportTimeout();

        if (_resultManager[reportId].tdsContractId != 0)
            revert TDSContractAlreadyReported();

        _;
    }

    /// -----------------------------------------------------------------------
    /// View Functions
    /// -----------------------------------------------------------------------

    function isTDSContractDefault(
        uint256 tdsContractId
    ) public view returns (bool) {
        return _tdsContractDefaultReportId[tdsContractId] != 0;
    }

    function getReportResultById(
        uint256 reportId
    ) public view returns (bool, Result memory) {
        return (
            _resultManager[reportId].tdsContractId != 0,
            _resultManager[reportId]
        );
    }

    function getReportIds(
        uint256 tdsContractId
    ) public view returns (uint256[] memory) {
        return _tdsContractReport[tdsContractId];
    }

    /// -----------------------------------------------------------------------
    /// Actions
    /// -----------------------------------------------------------------------

    function report(
        uint256 tdsContractId,
        uint32 referenceEvent,
        bytes calldata triggerType,
        bytes calldata proof
    ) external returns (uint256) {
        uint256 reportId = ++_lastReportId;
        _tdsContractReport[tdsContractId].push(reportId);
        uint64 startTime = uint64(block.timestamp);
        _reportManager[reportId] = Report({
            tdsContractId: tdsContractId,
            reportId: reportId,
            startTime: startTime,
            referenceEvent: referenceEvent,
            triggerType: triggerType,
            proof: proof
        });

        emit ReportCreated(
            reportId,
            tdsContractId,
            startTime,
            referenceEvent,
            triggerType,
            proof
        );

        return reportId;
    }

    function submitResult(
        uint32 reportId,
        bool isDefault,
        bytes calldata result
    ) external canSubmitResult(reportId) onlyOwner {
        uint256 tdsContractId = _reportManager[reportId].tdsContractId;
        uint64 submittedAt = uint64(block.timestamp);
        _resultManager[reportId] = Result({
            tdsContractId: tdsContractId,
            reportId: reportId,
            submittedAt: submittedAt,
            isDefault: isDefault,
            result: result
        });

        // Record first reportId confirm contract is default
        if (_tdsContractDefaultReportId[tdsContractId] == 0 && isDefault) {
            _tdsContractDefaultReportId[tdsContractId] = reportId;
        }

        emit ReportResultSubmitted(
            tdsContractId,
            reportId,
            submittedAt,
            isDefault,
            result
        );
    }
}
