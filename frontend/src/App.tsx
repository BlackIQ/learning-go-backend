import { useState, useEffect } from "react";

import { listTodos } from "./api/todo";
import type { TodoInterface } from "./types/todo";

function App() {
  const [loading, setLoading] = useState<boolean>(true);
  const [todos, setTodos] = useState<TodoInterface[]>([]);

  const [openCreate, setOpenCreate] = useState<boolean>(false);
  const handleCreate = () => setOpenCreate(!openCreate);

  const getTodos = async () => {
    try {
      const data = await listTodos();

      setTodos(data);
      setLoading(false);
    } catch {
      console.log("Issues");
    }
  };

  useEffect(() => {
    getTodos();
  }, []);

  return (
    <>
      <div className="container">
        <h1>Hello React!</h1>
        <p>Welcome to my new React application.</p>

        {!loading ? (
          <div className="container">
            <div className="row">
              <div className="col-3 mb-4">
                <div
                  style={{
                    display: "flex",
                    justifyContent: "center",
                    alignItems: "center",
                    textAlign: "center",
                    height: "100%",
                  }}
                >
                  <button className="btn btn-primary" onClick={handleCreate}>
                    New Todo
                  </button>
                </div>
              </div>
              {todos.map((todo) => (
                <div key={todo.id} className="col-3 mb-4">
                  <div className="card">
                    <div className="card-body">
                      <h4 className="card-title">{todo.title}</h4>
                      <p className="card-text">{todo.details}</p>
                      <button className="btn btn-primary btn-sm">Open</button>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        ) : (
          <div>
            <p>Loading todos</p>
          </div>
        )}

        {openCreate && (
          <>
            <div
              className="modal fade show d-block"
              tabIndex={-1}
              role="dialog"
            >
              <div className="modal-dialog modal-dialog-centered">
                <div className="modal-content">
                  <div className="modal-header">
                    <h5 className="modal-title">Create Todo</h5>
                    <button
                      type="button"
                      className="btn-close"
                      onClick={handleCreate}
                    />
                  </div>
                  <div className="modal-body">
                    <div>
                      <div className="mb-3">
                        <label
                          htmlFor="exampleFormControlInput1"
                          className="form-label"
                        >
                          Email address
                        </label>
                        <input
                          type="email"
                          className="form-control"
                          id="exampleFormControlInput1"
                          placeholder="name@example.com"
                        />
                      </div>
                      <div className="mb-3">
                        <label
                          htmlFor="exampleFormControlTextarea1"
                          className="form-label"
                        >
                          Example textarea
                        </label>
                        <textarea
                          className="form-control"
                          id="exampleFormControlTextarea1"
                          rows={3}
                        ></textarea>
                      </div>
                    </div>
                  </div>
                  <div className="modal-footer">
                    <button
                      className="btn btn-secondary"
                      onClick={handleCreate}
                    >
                      Close
                    </button>
                    <button className="btn btn-primary">Save changes</button>
                  </div>
                </div>
              </div>
            </div>

            <div className="modal-backdrop fade show" onClick={handleCreate} />
          </>
        )}
      </div>
    </>
  );
}

export default App;
