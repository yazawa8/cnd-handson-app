from sqlalchemy.orm import Session
from typing import Optional
from internal.role.model.role import RoleModel


class RoleRepository:
    def __init__(self, db: Session):
        """RoleRepositoryの初期化
        Args:
            db: SQLAlchemyのセッション
        """
        self.db = db

    def create(self, name: str, description: str) -> Optional[RoleModel]:
        """Roleの生成
        Args:
            name: Roleの名前
            description: Roleの説明
        """
        try:
            new_role = RoleModel(name=name, description=description)
            self.db.add(new_role)
            self.db.commit()
            self.db.refresh(new_role)
            return new_role
        except Exception as e:
            self.db.rollback()
            print(f"Error creating role: {e}")
            return None

    def get_by_id(self, Role_id: str) -> Optional[RoleModel]:
        """Roleの取得
        Args:
            Role_id: RoleのID
        """
        try:
            role = self.db.query(RoleModel).filter(RoleModel.id == Role_id).first()
            return role
        except Exception as e:
            print(f"Error fetching role by ID: {e}")
            return None

    def list_all(self) -> list[RoleModel]:
        """Roleの一覧取得"""
        try:
            role_list = self.db.query(RoleModel).all()
            return role_list
        except Exception as e:
            print(f"Error fetching all roles: {e}")
            return []

    def update(self, role: RoleModel) -> Optional[RoleModel]:
        """Roleの更新
        Args:
            role: Roleのインスタンス
        """
        try:
            self.db.add(role)
            self.db.commit()
            self.db.refresh(role)
            return role
        except Exception as e:
            self.db.rollback()
            print(f"Error updating role: {e}")
            return None

    def delete(self, role: RoleModel) -> bool:
        """Roleの削除
        Args:
            role: Roleのインスタンス
        """
        try:
            self.db.delete(role)
            self.db.commit()
            return True
        except Exception as e:
            self.db.rollback()
            print(f"Error deleting role: {e}")
            return False
