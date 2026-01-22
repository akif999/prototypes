import sys


def main():
    args = sys.argv[1:]

    if len(args) != 2 or args[0] != "--name":
        print("Usage: hello-python --name <name>")
        return

    name = args[1]
    print(f"Hello, {name}")


if __name__ == "__main__":
    main()
