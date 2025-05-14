from sqlalchemy.orm import Session
from typing import Optional
from internal.column.model.column import ColumnModel
from uuid import UUID


class ColumnRepository:
    def __init__(self, db: Session):
        """ColumnRepositoryの初期化
        Args:
            db: SQLAlchemyのセッション
        """
        self.db = db

    def create(self, name: str, board_id: UUID) -> Optional[ColumnModel]:
        """Columnの生成
        Args:
            name: Columnの名前
            board_id: BoardのID
        """
        try:
            new_column = ColumnModel(name=name, board_id=board_id)
            self.db.add(new_column)
            self.db.commit()
            self.db.refresh(new_column)
            return new_column
        except Exception as e:
            self.db.rollback()
            print(f"Error creating column: {e}")
            return None

    def get_by_id(self, column_id: str) -> Optional[ColumnModel]:
        """Columnの取得
        Args:
            column_id: ColumnのID
        """
        try:
            column = (
                self.db.query(ColumnModel).filter(ColumnModel.id == column_id).first()
            )
            return column
        except Exception as e:
            print(f"Error fetching column by ID: {e}")
            return None

    def update(self, column: ColumnModel) -> Optional[ColumnModel]:
        """Columnの更新
        Args:
            column: Columnのインスタンス
        """
        try:
            self.db.add(column)
            self.db.commit()
            self.db.refresh(column)
            return column
        except Exception as e:
            self.db.rollback()
            print(f"Error updating column: {e}")
            return None

    def delete(self, column: ColumnModel) -> bool:
        """Columnの削除
        Args:
            column: Columnのインスタンス
        """
        try:
            self.db.delete(column)
            self.db.commit()
            return True
        except Exception as e:
            self.db.rollback()
            print(f"Error deleting column: {e}")
            return False

    def list(
        self, board_id: UUID, page: int, page_size: int
    ) -> Optional[list[ColumnModel]]:
        """Columnの一覧取得
        Args:
            board_id: BoardのID
            page: ページ番号
            page_size: ページサイズ
        """
        try:
            columns = (
                self.db.query(ColumnModel)
                .filter(ColumnModel.board_id == board_id)
                .offset(max((page - 1), 0) * page_size)
                .limit(page_size)
                .all()
            )
            return columns
        except Exception as e:
            print(f"Error fetching columns: {e}")
            return None

    def total_count(self, board_id: UUID) -> int:
        """Columnの総数取得
        Args:
            board_id: BoardのID
        """
        try:
            total = (
                self.db.query(ColumnModel)
                .filter(ColumnModel.board_id == str(board_id))
                .count()
            )
            return total
        except Exception as e:
            print(f"Error fetching total count of columns: {e}")
            return 0
