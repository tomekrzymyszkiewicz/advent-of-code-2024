with open("input.txt", "r") as f:
    reports = [list(map(int, line.split())) for line in f.read().split("\n")]

SAFE_THRESHOLD_MAX = 3
SAFE_THRESHOLD_MIN = 1
safe_reports_count = 0


def is_report_safe(report: list) -> bool:
    return all(
        [
            SAFE_THRESHOLD_MAX >= abs(report[i] - report[i + 1]) >= SAFE_THRESHOLD_MIN
            for i in range(len(report) - 1)
        ]
    ) and (
        all(
            [report[i] > report[i + 1] for i in range(len(report) - 1)],
        )
        or all([report[i] < report[i + 1] for i in range(len(report) - 1)])
    )


for report in reports:
    if is_report_safe(report):
        safe_reports_count += 1
    else:
        for i in range(len(report)):
            if is_report_safe(report[:i] + report[i + 1 :]):
                safe_reports_count += 1
                break
print(safe_reports_count)
