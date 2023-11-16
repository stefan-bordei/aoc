import argparse, functools


def rgetattr(obj, attr, *args):
    def _getattr(obj, attr):
        return getattr(obj, attr, *args)
    return functools.reduce(_getattr, [obj] + attr.split('.'))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()

    parser.add_argument('--year', help='AoC Year. Mapped to src/y<year>.')
    parser.add_argument('--day', help='Day of AoC year.')
    
    args = parser.parse_args()
    
    mod = __import__(f'src.{args.year}.day{args.day}')
    runner = rgetattr(mod, f'{args.year}.day{args.day}.solve')
    runner()

