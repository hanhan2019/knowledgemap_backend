# -*- coding: utf-8 -*
import sys

def main():
    argc = len(sys.argv)
    # if argc<2:
    #     print('usage:%s <xxx.go>'%sys.argv[0])
    #     return
    f = open(sys.argv[1], 'rb')
    if not f:
        print("open file error:%s"%sys.argv[1])
        return
    s = f.read().decode('utf-8')
    f.close()
    s = s.replace(',omitempty', '')
    #
    f = open(sys.argv[1], 'wb')
    f.write(s.encode('utf-8'))
    f.close()
    print('ok')

if __name__=='__main__':
    main()