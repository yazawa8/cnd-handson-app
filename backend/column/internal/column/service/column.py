from internal.column.model.column import ColumnModel
from internal.column.repository.column import ColumnRepository
from uuid import UUID
from typing import Optional


class ColumnService:
    def __init__(self, repository: ColumnRepository) -> None:
        """ColumnServiceの初期化
        Args:
            repository: ColumnRepositoryのインスタンス
        """
        self.repository = repository

    def create(self, name: str, board_id: UUID) -> Optional[ColumnModel]:
        """Columnの生成
        Args:
            name: Columnの名前
            board_id: BoardのID
        """
        return self.repository.create(name, board_id)

    def get(self, column_id: UUID) -> Optional[ColumnModel]:
        """Columnの取得
        Args:
            column_id: ColumnのID
        """
        return self.repository.get_by_id(column_id)

    def update(
        self,
        column_id: UUID,
        name: Optional[str] = None,
        board_id: Optional[str] = None,
    ) -> Optional[ColumnModel]:
        """Columnの更新
        Args:
            column_id: ColumnのID
            name: Columnの名前
            description: Columnの説明
        """
        if name is None and board_id is None:
            raise ValueError("At least one of name or description must be provided")

        column = self.repository.get_by_id(column_id)
        if name is not None or name != "":
            column.name = name
        if board_id is not None:
            column.description = board_id

        self.repository.update(column)
        return column

    def delete(self, column_id: UUID) -> bool:
        """Columnの削除
        Args:
            column_id: ColumnのID
        """
        column = self.repository.get_by_id(column_id)
        if column is None:
            raise ValueError("Column not found")
        if column.name == "admin":
            raise ValueError("Cannot delete admin column")

        return self.repository.delete(column)

    def list(
        self, board_id: UUID, page: int, page_size: int
    ) -> Optional[list[ColumnModel]]:
        """Columnの一覧取得
        Args:
            board_id: BoardのID
            page: ページ番号
            page_size: ページサイズ
        """
        return self.repository.list(board_id, page, page_size)

    def total_count(self, board_id: UUID) -> int:
        """Columnの総数取得
        Args:
            board_id: BoardのID
        """
        return self.repository.total_count(board_id)
