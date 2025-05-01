from proto import role_pb2_grpc, role_pb2

class RoleHandler(role_pb2_grpc.RoleServiceServicer):
    """RoleHandlerの実装
    """
    def __init__(self,role_service):
        """RoleHandlerの初期化
        Args:
            role_service: RoleServiceのインスタンス
        """
        self.role_service = role_service

    def CreateRole(self, request, context):
        """Roleの生成
        Args:
        request: CreateRoleRequest
            Roleの生成に必要な情報を含むリクエスト
        context: grpc.ServicerContext
        """
        role = request.role
        # ここでRoleを生成する処理を実装する
        role = self.role_service.create(role.name, role.description)
        response = role_pb2.RoleResponse(role=role)
        return response
    def GetRole(self, request, context):
        """Roleの取得
        """
        role_id = request.role_id
        # ここでRoleを取得する処理を実装する
        role = self.role_service.get(id=role_id, name="example_role")
        response = role_pb2.RoleResponse(role=role)
        return response
    def UpdateRole(self, request, context):
        """Roleの更新
        """
        role = request.role
        # ここでRoleを更新する処理を実装する
        role = self.role_service.update(role.id, role.name, role.description)
        if role is None:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details('Role not found!')
            return role_pb2.RoleResponse()
        response = role_pb2.RoleResponse(role=role)
        return response
    def DeleteRole(self, request, context):
        """Roleの削除
        """
        role_id = request.role_id
        # ここでRoleを削除する処理を実装する
        success = self.role_service.delete(role_id)
        if not success:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details('Role not found!')
            return role_pb2.DeleteRoleResponse()
        response = role_pb2.DeleteRoleResponse(role_id=role_id)
        return response
    def ListRoles(self, request, context):
        """Roleの一覧取得
        """
        # ここでRoleの一覧を取得する処理を実装する
        roles = [role_pb2.Role(id="1", name="example_role_1"),
                 role_pb2.Role(id="2", name="example_role_2")]
        response = role_pb2.ListRolesResponse(roles=roles)
        return response
    
