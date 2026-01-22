import argparse
import sys
from pathlib import Path

from hello_python.counter import (
    count_names,
    sort_counts,
    InvalidCSVError,
)
from hello_python.formatter import format_plain, format_table


def main() -> None:
    parser = argparse.ArgumentParser(
        prog="hello-python",
        description="Count names from csv",
    )

    parser.add_argument(
        "--input",
        required=True,
        help="input csv file",
    )
    parser.add_argument(
        "--sort",
        choices=["name", "count"],
        default="name",
    )
    parser.add_argument(
        "--format",
        choices=["plain", "table"],
        default="plain",
    )

    args = parser.parse_args()

    try:
        counts = count_names(Path(args.input))
        items = sort_counts(counts, by=args.sort)
    except (FileNotFoundError, InvalidCSVError, ValueError) as e:
        print(e, file=sys.stderr)
        sys.exit(1)

    if args.format == "plain":
        output = format_plain(items)
    else:
        output = format_table(items)
    print(output)


if __name__ == "__main__":
    main()
