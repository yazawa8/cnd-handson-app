import grpc
from proto import role_pb2_grpc, role_pb2
import argparse  # argparseをインポート
import os
from typing import Any

def create_Request(command: str, arguments: list[str]) -> Any:
    """コマンドに応じたリクエストを生成する関数
    Args:
        command: コマンド名
    """
    if command == "create":
        return role_pb2_grpc.CreateRoleRequest(
            role=role_pb2.Role(
                id = arguments[0],
                name = arguments[1],
                description = arguments[2]
            )
        )
    elif command == "get":
        return role_pb2_grpc.GetRoleRequest(role_id=arguments[0])
    elif command == "update":
        return role_pb2_grpc.UpdateRoleRequest(
            role=role_pb2.Role(
                id = arguments[0],
                name = arguments[1],
                description = arguments[2]
            )
        )
    elif command == "delete":
        return role_pb2_grpc.DeleteRoleRequest(role_id=arguments[0])
    elif command == "list":
        return role_pb2_grpc.ListRolesRequest()
    else:
        raise ValueError("Invalid command")
def run(server_address: str, command: str, arguments: list[str]):
    # サーバーのアドレスとポートを指定

    # gRPCチャンネルを作成
    with grpc.insecure_channel(server_address) as channel:
        # スタブを作成
        stub = role_pb2_grpc.RoleServiceStub(channel)
        # コマンドに応じたリクエストを生成
        request = create_Request(command, arguments)
        response = stub.YourMethod(request)

        # レスポンスを表示
        print("Response from server:", response.message)

if __name__ == "__main__":

    # コマンドライン引数のパーサーを作成
    parser = argparse.ArgumentParser(description='gRPC Client for Role Service')
    parser.add_argument('--port', type=int, default=50051, help='Port number to connect to the server')
    args = parser.parse_args()
    # ポート番号を取得
    port = args.port
    # サーバーのアドレスを指定
    server_address = f'localhost:{port}'

    # コマンドライン引数を取得
    command = input("Enter command (create, get, update, delete, list): ").strip()
    arguments = input("Enter arguments (comma separated): ").strip().split(",")
    # 引数をトリミング
    arguments = [arg.strip() for arg in arguments]
    # コマンドを実行
    try:
        run(server_address, command, arguments)
    except Exception as e:
        print(f"Error: {e}")
