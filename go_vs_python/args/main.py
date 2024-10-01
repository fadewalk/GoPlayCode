import sys

def tansform(*args):
    for arg in args:
        print(arg.upper())

if __name__ == "__main__":
    tansform(*sys.argv[1:])