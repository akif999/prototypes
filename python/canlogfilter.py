import re
import argparse

def hexint(x):
    return int(x, 16)

class Record:
    def __init__(self):
        self.crtime   = 0.00000
        self.ch       = 1
        self.hexid    = 0x000
        self.dir      = "Rx"
        self.stat     = "d"
        self.dlc      = 0
        self.data     = []
        self.length   = 0
        self.bitcount = 0
        self.decid    = 0

def main():
    parser = argparse.ArgumentParser(description = 'CanlogFilter')

    parser.add_argument('inputFile',        help = 'Input file path')
    parser.add_argument('--difftime', '-d', action = 'store_const', const = True, default = False,  help = 'Print with difftime')
    parser.add_argument('--pickup',   '-u', nargs = '*', help = 'pick up records')
    parser.add_argument('--dropoff',  '-o', nargs = '*', help = 'drop off records')
    args = parser.parse_args()

    canlog = []
    canlog = parse(args.inputFile)

    if args.pickup != None and args.dropoff != None:
        print "--pickup and --dropoff, both provide"
        return -1
    elif args.pickup != None:
        canlog = pick_log(canlog, map(hexint, args.pickup))
    elif args.dropoff != None:
        canlog = drop_log(canlog, map(hexint, args.dropoff))

    if args.difftime == True:
        printlog_with_diff_time(canlog)
    else:
        printlog(canlog)

def parse(filename):
    canlog = []
    for line in open(filename, 'r'):
        fields = line.split()
        if re.match("1|2", fields[1]):
            rec = Record()
            rec.crtime   = float(fields[0])
            rec.ch       = int(fields[1], 10)
            rec.hexid    = int(fields[2], 16)
            rec.dir      = fields[3]
            rec.stat     = fields[4]
            rec.dlc      = int(fields[5], 10)
            rec.data     = map(hexint, fields[6:rec.dlc+6])
            rec.length   = int(fields[rec.dlc+8], 10)
            rec.bitcount = int(fields[rec.dlc+11], 10)
            rec.decid    = int(fields[rec.dlc+14], 10)
            canlog.append(rec)
    return canlog

def pick_log(canlog, ids):
    ret = []
    for rec in canlog:
        if rec.hexid in ids:
            ret.append(rec)
    return ret

def drop_log(canlog, ids):
    ret = []
    for rec in canlog:
        if not rec.hexid in ids:
            ret.append(rec)
    return ret

def printlog(canlog):
    for rec in canlog:
        print '%f %d %03X %s %d' % (rec.crtime, rec.ch, rec.hexid, rec.dir, rec.dlc),
        for byte in rec.data:
            print '%02X' % byte,
        print

def printlog_with_diff_time(canlog):
    prevtime = 0
    difftime = 0
    for rec in canlog:
        difftime = rec.crtime - prevtime
        print '%f %f %d %03X %s %d' % (difftime, rec.crtime, rec.ch, rec.hexid, rec.dir, rec.dlc),
        for byte in rec.data:
            print '%02X' % byte,
        print
        prevtime = rec.crtime

if __name__ == '__main__' :
    main()
