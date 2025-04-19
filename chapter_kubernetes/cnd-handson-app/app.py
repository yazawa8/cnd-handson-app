from flask import Flask, render_template
import logging
import socket
import mysql.connector
import os

app = Flask(__name__)

DB_Host = os.environ.get('DB_Host') or "localhost"
DB_Database = os.environ.get('DB_Database') or "mysql"
DB_User = os.environ.get('DB_User') or "root"
DB_Password = os.environ.get('DB_Password') or "cnds2025"


logging.basicConfig(level=logging.INFO,
                    format='%(asctime)s [%(levelname)s] - %(message)s',
                    handlers=[logging.StreamHandler()])

def check_db_connection():
    try:
        logging.info("Attempting to connect to the database...")
        connection = mysql.connector.connect(
            host=DB_Host,
            database=DB_Database,
            user=DB_User,
            password=DB_Password
        )
        connection.close()
        logging.info("Database connection successful.")
        return True, ""
    except Exception as e:
        logging.error("Database connection failed: %s", str(e))
        return False, str(e)

@app.route("/")
def main():
    db_connect_result, err_message = check_db_connection()
    color = '#bfbed3' if db_connect_result else '#c53d43'
    
    if err_message:
        logging.debug("Error message: %s", err_message)

    return render_template('index.html', debug=err_message, db_connect_result=db_connect_result, name=socket.gethostname(), color=color)

@app.route("/debug")
def debug():
    color = '#2196f3'
    debug_info = (
        "Environment Variables: DB_Host=" + (os.environ.get('DB_Host') or "Not Set") +
        "; DB_Database=" + (os.environ.get('DB_Database') or "Not Set") +
        "; DB_User=" + (os.environ.get('DB_User') or "Not Set") +
        "; DB_Password=" + (os.environ.get('DB_Password') or "Not Set")
    )
    logging.debug(debug_info)
    return render_template('index.html', debug=debug_info, color=color)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8888)
