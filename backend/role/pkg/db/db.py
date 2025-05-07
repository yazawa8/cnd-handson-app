from sqlalchemy import create_engine, Engine
from sqlalchemy.orm import sessionmaker, Session
from contextlib import contextmanager


class Database:
    def __init__(self, db_url: str, models: list) -> None:
        """Databaseの初期化
        Args:
            db_url: データベースのURL
            models: SQLAlchemyのモデルクラスのリスト
        """
        self.engine = create_engine(db_url, echo=True)  # echo=TrueはSQLログ出したいなら
        self.SessionLocal = sessionmaker(bind=self.engine, autoflush=False, autocommit=False)
        self.models = models

    def get_engine(self) -> Engine:
        """SQLAlchemyエンジンの取得
        """
        return self.engine
    
    def get_session(self) -> Session:
        """SQLAlchemyセッションの取得
        """
        return self.SessionLocal()
    
    def migrate(self) -> None:
        """データベースのマイグレーション
        """
        for model in self.models:
            model.__table__.create(self.engine, checkfirst=True)

        return None
    
    def drop_all(self) -> None:
        """データベースの全テーブル削除
        """
        for model in self.models:
            model.__table__.drop(self.engine, checkfirst=True)

        return None

    def close(self) -> None:
        """データベース接続のクローズ
        """
        self.engine.dispose()
        return None

    @contextmanager
    def session_scope(self) -> Session:
        """with構文でセッション管理する"""
        session = self.get_session()
        try:
            yield session
            session.commit()
        except Exception:
            session.rollback()
            raise
        finally:
            session.close()
