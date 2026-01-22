import argparse


def main():
    parser = argparse.ArgumentParser(
        prog="hello-python",
        description="Simple greeting CLI",
    )

    parser.add_argument(
        "--name",
        required=True,
        help="name to greet",
    )

    args = parser.parse_args()

    print(f"Hello, {args.name}")


if __name__ == "__main__":
    main()
