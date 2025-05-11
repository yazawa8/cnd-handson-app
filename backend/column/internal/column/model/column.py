from sqlalchemy import Column, VARCHAR
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.ext.declarative import declarative_base
import uuid

Base = declarative_base()


class ColumnModel(Base):
    """Columnモデル
    Args:
        id: UUID RoleのID
        name: str Roleの名前
        board_id: UUID Roleの説明
    """
    __tablename__ = 'roles'

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4, index=True)
    name = Column(VARCHAR(255), nullable=False, index=True)
    board_id = Column(UUID(as_uuid=True), default=uuid.uuid4, index=True)
