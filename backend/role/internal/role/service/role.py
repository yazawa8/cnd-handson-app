from internal.role.model.role import RoleModel
from internal.role.repository.role import RoleRepository
from typing import Optional
from uuid import UUID


class RoleService:
    def __init__(self, repository: RoleRepository) -> None:
        """RoleServiceの初期化
        Args:
            repository: RoleRepositoryのインスタンス
        """
        self.repository = repository

    def create(self, name: str, description: str) -> Optional[RoleModel]:
        """Roleの生成
        Args:
            name: Roleの名前
            description: Roleの説明
        """
        return self.repository.create(name, description)

    def get(self, role_id: UUID) -> Optional[RoleModel]:
        """Roleの取得
        Args:
            role_id: RoleのID
        """
        return self.repository.get_by_id(role_id)

    def update(
        self,
        role_id: UUID,
        name: Optional[str] = None,
        description: Optional[str] = None,
    ) -> Optional[RoleModel]:
        """Roleの更新
        Args:
            role_id: RoleのID
            name: Roleの名前
            description: Roleの説明
        """
        if name is None and description is None:
            raise ValueError("At least one of name or description must be provided")
        if name == "":
            raise ValueError("Name cannot be empty")

        role = self.repository.get_by_id(role_id)
        if role is not None:
            role.name = name
        if description is not None:
            role.description = description

        self.repository.update(role)
        return role

    def delete(self, role_id: UUID) -> bool:
        """Roleの削除
        Args:
            role_id: RoleのID
        """
        role = self.repository.get_by_id(role_id)
        if role is None:
            raise ValueError("Role not found")
        if role.name == "admin":
            raise ValueError("Cannot delete admin role")

        return self.repository.delete(role)

    def list(self, page: int, page_size: int) -> Optional[list[RoleModel]]:
        """Roleの一覧取得
        Args:
            page: ページ番号
            page_size: ページサイズ
        """
        return self.repository.find_all(page, page_size)
