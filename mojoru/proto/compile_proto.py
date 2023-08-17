

# 介绍python获取命令行参数的方法：getopt模和argparse模块。
import os
import sys
import getopt
import subprocess


def main(argv):
    """

    python3 compile_proto.py -p  XXXX.proto
    :param argv:
    :return:
    """
    protofile = ""

    try:
        options, args = getopt.getopt(argv, "hp:", ['help', "protofile="])

        for option, value in options:
            if option in ("-p", "--protofile"):
                protofile = value

# gRPC  +  micro
# protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
        absolutePath = os.getcwd()
        protofilePath = os.path.join(absolutePath, protofile)
        subprocess.call("export PATH=$PATH:/usr/local/go/bin;", shell=True)
        subprocess.call(
            "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2", shell=True)
        subprocess.call(
            "go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28", shell=True)
        subprocess.call(
            "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest", shell=True)
        subprocess.call(
            "go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest", shell=True)
        subprocess.call(
            "go install github.com/micro/micro/v3@latest", shell=True)
        subprocess.call(
            "go install go-micro.dev/v4@latest", shell=True)
        # --go_out: protoc-gen-go: Plugin failed with status code 1.の解決方法
        subprocess.call("export GOPATH=$HOME/go", shell=True)
        subprocess.call("PATH=$PATH:$GOPATH/bin", shell=True)
        # subprocess.call('export PATH="$PATH:$(go env GOPATH)/bin')

        #compile_cmd  = " protoc -I={1} --go_out={1} {0}".format(protofilePath,absolutePath)
        compile_cmd_grpc = "protoc --go_out=. --go-grpc_out=. {0}".format(
            protofile)
        compile_cmd_micro = "protoc --proto_path=. --micro_out=. --go_out=. {0}".format(
            protofile)
        subprocess.call(compile_cmd_grpc, shell=True)
        subprocess.call(compile_cmd_micro, shell=True)
        print(compile_cmd_grpc)
        print(compile_cmd_micro)

    except getopt.GetoptError:
        print("###################################################")
        print("\n")
        print("\n")
        print("\n")
        print("python3 compile_proto.py -p  XXXX.proto")
        print("\n")
        print("\n")
        print("\n")
        print("###################################################")
        sys.exit()


if __name__ == '__main__':
    main(sys.argv[1:])
