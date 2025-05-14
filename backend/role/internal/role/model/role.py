from sqlalchemy import Column, VARCHAR, TEXT
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.ext.declarative import declarative_base
import uuid

Base = declarative_base()


class RoleModel(Base):
    """Roleモデル
    Args:
        id: UUID RoleのID
        name: str Roleの名前
        description: str Roleの説明
    """

    __tablename__ = "roles"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4, index=True)
    name = Column(VARCHAR(255), nullable=False, index=True)
    description = Column(TEXT, index=True)
