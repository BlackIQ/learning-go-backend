import { useState, useEffect } from "react";

import { listTodos } from "./api/todo";

function App() {
  const [todos, setTodos] = useState([]);

  const getTodos = async () => {
    try {
      const data = await listTodos();

      console.log(data);
    } catch {
      console.log("Issues");
    }
  };

  useEffect(() => {
    getTodos();
  }, []);

  return (
    <div>
      <h1>Hello React!</h1>
      <p>Welcome to my new React application.</p>

      <button
        type="button"
        className="btn btn-primary"
        onClick={() => {
          console.log("Bitch clicked!");
        }}
      >
        Click on me bitch!
      </button>
    </div>
  );
}

export default App;
