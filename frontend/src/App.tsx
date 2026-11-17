import { useState, useEffect } from "react";
import { useForm, type SubmitHandler } from "react-hook-form";

import { listTodos, createTodo, updateTodo, deleteTodo } from "./api/todo";
import type {
  TodoInterface,
  CreateTodoInterface,
  UpdateTodoInterface,
} from "./types/todo";

function App() {
  const [loading, setLoading] = useState<boolean>(true);

  const [todos, setTodos] = useState<TodoInterface[]>([]);
  const [currentTodo, setCurrentTodo] = useState<TodoInterface>();

  const [createLoading, setCreateLoading] = useState<boolean>(false);
  const [openCreate, setOpenCreate] = useState<boolean>(false);
  const handleCreateModal = () => setOpenCreate(!openCreate);

  const [updateLoading, setUpdateLoading] = useState<boolean>(false);
  const [openUpdate, setOpenUpdate] = useState<boolean>(false);
  const handleUpdateModal = () => setOpenUpdate(!openUpdate);

  const {
    register: registerCreate,
    handleSubmit: handleSubmitCreate,
    reset: resetCreate,
    formState: { errors: errorsCreate },
  } = useForm<CreateTodoInterface>();

  const {
    register: registerUpdate,
    handleSubmit: handleSubmitUpdate,
    reset: resetUpdate,
    formState: { errors: errorsUpdate },
  } = useForm<UpdateTodoInterface>();

  const getTodos = async () => {
    try {
      const data = await listTodos();

      setTodos(data);
      setLoading(false);
    } catch {
      console.log("Issues");
    }
  };

  const onSubmitCreate: SubmitHandler<CreateTodoInterface> = async (data) => {
    setCreateLoading(true);

    try {
      await createTodo(data);

      resetCreate({});

      handleCreateModal();

      setCreateLoading(false);

      getTodos();
    } catch {
      console.error("Error in creating data.");
      setCreateLoading(false);
    }
  };

  const onSubmitUpdate: SubmitHandler<UpdateTodoInterface> = async (data) => {
    setUpdateLoading(true);

    if (currentTodo) {
      try {
        await updateTodo(currentTodo.id, data);

        resetUpdate({});

        handleUpdateModal();

        setUpdateLoading(false);

        getTodos();
      } catch {
        console.error("Error in updating data.");
        setUpdateLoading(false);
      }
    }
  };

  const onSubmitDelete = async (id: string) => {
    try {
      await deleteTodo(id);

      getTodos();
    } catch {
      console.error("Error in updating data.");
    }
  };

  useEffect(() => {
    getTodos();
  }, []);

  useEffect(() => {
    if (currentTodo) {
      resetUpdate({
        title: currentTodo.title,
        details: currentTodo.details,
      });
    }
  }, [currentTodo, resetUpdate]);

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
                  <button
                    className="btn btn-primary"
                    onClick={handleCreateModal}
                  >
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
                      <button
                        className="btn btn-sm btn-primary"
                        onClick={() => {
                          setCurrentTodo(todo);
                          handleUpdateModal();
                        }}
                      >
                        Open
                      </button>
                      <button
                        className="btn btn-sm btn-danger"
                        onClick={() => onSubmitDelete(todo.id)}
                      >
                        Delete
                      </button>
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
          <form onSubmit={handleSubmitCreate(onSubmitCreate)}>
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
                      onClick={handleCreateModal}
                    />
                  </div>
                  <div className="modal-body">
                    <div>
                      <div className="mb-3">
                        <label htmlFor="title" className="form-label">
                          Title
                        </label>
                        <input
                          id="title"
                          type="text"
                          className="form-control"
                          placeholder="Enter title"
                          {...registerCreate("title", {
                            required: "Title is required",
                            disabled: createLoading,
                          })}
                        />
                        <span
                          className={
                            errorsCreate.title ? "text-danger" : "text-success"
                          }
                        >
                          {errorsCreate.title
                            ? errorsCreate.title.message
                            : "Looks good!"}
                        </span>
                      </div>
                      <div className="mb-3">
                        <label htmlFor="details" className="form-label">
                          Details
                        </label>
                        <textarea
                          id="details"
                          className="form-control"
                          placeholder="Enter Details"
                          {...registerCreate("details", {
                            required: "Details is required",
                            disabled: createLoading,
                          })}
                          rows={3}
                        />
                        <span
                          className={
                            errorsCreate.details
                              ? "text-danger"
                              : "text-success"
                          }
                        >
                          {errorsCreate.details
                            ? errorsCreate.details.message
                            : "Looks good!"}
                        </span>
                      </div>
                    </div>
                  </div>
                  <div className="modal-footer">
                    <button
                      type="button"
                      className="btn btn-secondary"
                      onClick={handleCreateModal}
                    >
                      Close
                    </button>
                    <button
                      type="submit"
                      className="btn btn-primary"
                      disabled={createLoading}
                    >
                      Create Todo
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <div
              className="modal-backdrop fade show"
              onClick={handleCreateModal}
            />
          </form>
        )}

        {openUpdate && (
          <form onSubmit={handleSubmitUpdate(onSubmitUpdate)}>
            <div
              className="modal fade show d-block"
              tabIndex={-1}
              role="dialog"
            >
              <div className="modal-dialog modal-dialog-centered">
                <div className="modal-content">
                  <div className="modal-header">
                    <h5 className="modal-title">Update Todo</h5>
                    <button
                      type="button"
                      className="btn-close"
                      onClick={handleUpdateModal}
                    />
                  </div>
                  <div className="modal-body">
                    <div>
                      <div className="mb-3">
                        <label htmlFor="title" className="form-label">
                          Title
                        </label>
                        <input
                          id="title"
                          type="text"
                          className="form-control"
                          placeholder="Enter title"
                          {...registerUpdate("title", {
                            required: "Title is required",
                            disabled: updateLoading,
                          })}
                        />
                        <span
                          className={
                            errorsUpdate.title ? "text-danger" : "text-success"
                          }
                        >
                          {errorsUpdate.title
                            ? errorsUpdate.title.message
                            : "Looks good!"}
                        </span>
                      </div>
                      <div className="mb-3">
                        <label htmlFor="details" className="form-label">
                          Details
                        </label>
                        <textarea
                          id="details"
                          className="form-control"
                          placeholder="Enter Details"
                          {...registerUpdate("details", {
                            required: "Details is required",
                            disabled: updateLoading,
                          })}
                          rows={3}
                        />
                        <span
                          className={
                            errorsUpdate.details
                              ? "text-danger"
                              : "text-success"
                          }
                        >
                          {errorsUpdate.details
                            ? errorsUpdate.details.message
                            : "Looks good!"}
                        </span>
                      </div>
                    </div>
                  </div>
                  <div className="modal-footer">
                    <button
                      type="button"
                      className="btn btn-secondary"
                      onClick={handleUpdateModal}
                    >
                      Close
                    </button>
                    <button
                      type="submit"
                      className="btn btn-primary"
                      disabled={updateLoading}
                    >
                      Update Todo
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <div
              className="modal-backdrop fade show"
              onClick={handleUpdateModal}
            />
          </form>
        )}
      </div>
    </>
  );
}

export default App;
