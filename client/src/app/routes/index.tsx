// src/app/routes/index.tsx
import React from "react";
import { useAuth0 } from "@auth0/auth0-react";

const Home: React.FC = () => {
  const { loginWithRedirect, isAuthenticated, logout } = useAuth0();

  return (
    <div style={styles.container}>
      <h1 style={styles.title}>Todoアプリへようこそ</h1>
      {!isAuthenticated ? (
        <div style={styles.buttonContainer}>
          <button style={styles.button} onClick={() => loginWithRedirect()}>
            ログイン
          </button>
          {/*<button*/}
          {/*  style={styles.button}*/}
          {/*  onClick={() =>*/}
          {/*    loginWithRedirect()*/}
          {/*  }*/}
          {/*>*/}
          {/*  サインアップ*/}
          {/*</button>*/}
        </div>
      ) : (
        <div style={styles.buttonContainer}>
          <button style={styles.button} onClick={() => logout({logoutParams: {returnTo: window.location.origin} })}>
            ログアウト
          </button>
        </div>
      )}
    </div>
  );
};

// スタイルオブジェクト（シンプルなインラインスタイルを使用）
const styles: { [key: string]: React.CSSProperties } = {
  container: {
    textAlign: "center",
    marginTop: "50px",
  },
  title: {
    fontSize: "2.5rem",
    marginBottom: "30px",
  },
  buttonContainer: {
    display: "flex",
    justifyContent: "center",
    gap: "20px",
  },
  button: {
    padding: "10px 20px",
    fontSize: "1rem",
    cursor: "pointer",
    border: "none",
    borderRadius: "5px",
    backgroundColor: "#007bff",
    color: "#fff",
    transition: "background-color 0.3s",
  },
};

// ボタンのホバー効果を追加するためのCSS（オプション）
const styleSheet = document.createElement("style");
styleSheet.type = "text/css";
styleSheet.innerText = `
  button:hover {
    background-color: #0056b3;
  }
`;
document.head.appendChild(styleSheet);

export default Home;
