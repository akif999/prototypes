import argparse
from pathlib import Path

from hello_python.counter import count_names


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

    args = parser.parse_args()
    input_path = Path(args.input)

    counts = count_names(input_path)

    for name, count in counts.items():
        print(f"{name}: {count}")


if __name__ == "__main__":
    main()
