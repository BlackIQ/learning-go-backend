import { useState, useEffect } from "react";
import { useForm, type SubmitHandler } from "react-hook-form";

import { listTodos, createTodo, updateTodo, deleteTodo } from "./api/todo";
import type {
  TodoInterface,
  CreateTodoInterface,
  UpdateTodoInterface,
} from "./types/todo";

function App() {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [todos, setTodos] = useState<TodoInterface[]>([]);
  const [currentTodo, setCurrentTodo] = useState<TodoInterface | null>(null);

  // Create modal
  const [createLoading, setCreateLoading] = useState(false);
  const [openCreate, setOpenCreate] = useState(false);

  // Update modal
  const [updateLoading, setUpdateLoading] = useState(false);
  const [openUpdate, setOpenUpdate] = useState(false);

  const {
    register: registerCreate,
    handleSubmit: handleSubmitCreate,
    reset: resetCreate,
    formState: { errors: errorsCreate, dirtyFields: dirtyCreate },
  } = useForm<CreateTodoInterface>();

  const {
    register: registerUpdate,
    handleSubmit: handleSubmitUpdate,
    reset: resetUpdate,
    formState: { errors: errorsUpdate, dirtyFields: dirtyUpdate },
  } = useForm<UpdateTodoInterface>();

  const getTodos = async () => {
    setError(null);
    try {
      const data = await listTodos();
      setTodos(data);
    } catch {
      setError("Failed to fetch todos");
      console.error("Failed to fetch todos");
    } finally {
      setLoading(false);
    }
  };

  const onSubmitCreate: SubmitHandler<CreateTodoInterface> = async (data) => {
    setCreateLoading(true);
    setError(null);

    try {
      await createTodo(data);
      resetCreate();
      setOpenCreate(false);
      await getTodos();
    } catch {
      setError("Error creating todo");
      console.error("Error creating todo");
    } finally {
      setCreateLoading(false);
    }
  };

  const onSubmitUpdate: SubmitHandler<UpdateTodoInterface> = async (data) => {
    if (!currentTodo) return;

    setUpdateLoading(true);
    setError(null);

    try {
      await updateTodo(currentTodo.id, data);
      resetUpdate();
      setOpenUpdate(false);
      setCurrentTodo(null);
      await getTodos();
    } catch {
      setError("Error updating todo");
      console.error("Error updating todo");
    } finally {
      setUpdateLoading(false);
    }
  };

  const handleDelete = async (id: string) => {
    const confirmed = window.confirm("Are you sure?");
    if (!confirmed) return;

    try {
      await deleteTodo(id);
      await getTodos();
    } catch {
      setError("Error deleting todo");
      console.error("Error deleting todo");
    }
  };

  const openUpdateModal = (todo: TodoInterface) => {
    setCurrentTodo(todo);
    resetUpdate({
      title: todo.title,
      details: todo.details,
    });
    setOpenUpdate(true);
  };

  useEffect(() => {
    getTodos();
  }, []);

  const renderFieldError = (
    error: { message?: string } | undefined,
    isDirty: boolean | undefined,
  ) => {
    if (error) {
      return <span className="text-danger small">{error.message}</span>;
    }
    if (isDirty) {
      return <span className="text-success small">Looks good!</span>;
    }
    return null;
  };

  return (
    <div className="container py-4">
      <h1 className="mb-2">Gin and React</h1>
      <p className="text-muted mb-4">
        If you are at this page it means you executed the binary file generated
        with <b>Gin</b>.
      </p>

      {error && (
        <div className="alert alert-danger d-flex justify-content-between align-items-center">
          <span>{error}</span>
          <button
            type="button"
            className="btn-close"
            onClick={() => setError(null)}
          />
        </div>
      )}

      {loading ? (
        <div className="text-center py-5">
          <div className="spinner-border text-primary" role="status">
            <span className="visually-hidden">Loading...</span>
          </div>
          <p className="mt-2">Loading Todos...</p>
        </div>
      ) : (
        <>
          <div className="mb-4">
            <button
              className="btn btn-primary"
              onClick={() => setOpenCreate(true)}
            >
              + New Todo
            </button>
          </div>

          {todos.length === 0 ? (
            <div className="text-center text-muted py-5">No Todos yet!</div>
          ) : (
            <div className="row g-3">
              {todos.map((todo) => (
                <div
                  key={todo.id}
                  className="col-12 col-md-6 col-lg-4 col-xl-3"
                >
                  <div className="card h-100 shadow-sm">
                    <div className="card-body d-flex flex-column">
                      <h5 className="card-title">{todo.title}</h5>
                      <p className="card-text flex-grow-1">{todo.details}</p>
                      <div className="d-flex gap-2 mt-auto">
                        <button
                          className="btn btn-sm btn-outline-primary"
                          onClick={() => openUpdateModal(todo)}
                        >
                          Edit
                        </button>
                        <button
                          className="btn btn-sm btn-outline-danger"
                          onClick={() => handleDelete(todo.id)}
                        >
                          Delete
                        </button>
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          )}
        </>
      )}

      {/* ==================== Create Modal ==================== */}
      {openCreate && (
        <>
          <div
            className="modal fade show d-block"
            tabIndex={-1}
            role="dialog"
            aria-modal="true"
          >
            <div className="modal-dialog modal-dialog-centered">
              <form
                className="modal-content"
                onSubmit={handleSubmitCreate(onSubmitCreate)}
              >
                <div className="modal-header">
                  <h5 className="modal-title">Create Todo</h5>
                  <button
                    type="button"
                    className="btn-close"
                    onClick={() => setOpenCreate(false)}
                    disabled={createLoading}
                  />
                </div>

                <div className="modal-body">
                  <div className="mb-3">
                    <label htmlFor="create-title" className="form-label">
                      Title
                    </label>
                    <input
                      id="create-title"
                      type="text"
                      className={`form-control ${
                        errorsCreate.title
                          ? "is-invalid"
                          : dirtyCreate.title
                            ? "is-valid"
                            : ""
                      }`}
                      placeholder="Enter title"
                      disabled={createLoading}
                      {...registerCreate("title", {
                        required: "Title is required",
                      })}
                    />
                    {renderFieldError(errorsCreate.title, dirtyCreate.title)}
                  </div>

                  <div className="mb-3">
                    <label htmlFor="create-details" className="form-label">
                      Details
                    </label>
                    <textarea
                      id="create-details"
                      className={`form-control ${
                        errorsCreate.details
                          ? "is-invalid"
                          : dirtyCreate.details
                            ? "is-valid"
                            : ""
                      }`}
                      placeholder="Enter details"
                      rows={3}
                      disabled={createLoading}
                      {...registerCreate("details", {
                        required: "Details is required",
                      })}
                    />
                    {renderFieldError(
                      errorsCreate.details,
                      dirtyCreate.details,
                    )}
                  </div>
                </div>

                <div className="modal-footer">
                  <button
                    type="button"
                    className="btn btn-secondary"
                    onClick={() => setOpenCreate(false)}
                    disabled={createLoading}
                  >
                    Close
                  </button>
                  <button
                    type="submit"
                    className="btn btn-primary"
                    disabled={createLoading}
                  >
                    {createLoading ? (
                      <>
                        <span
                          className="spinner-border spinner-border-sm me-2"
                          role="status"
                          aria-hidden="true"
                        />
                        Creating...
                      </>
                    ) : (
                      "Create Todo"
                    )}
                  </button>
                </div>
              </form>
            </div>
          </div>
          <div
            className="modal-backdrop fade show"
            onClick={() => !createLoading && setOpenCreate(false)}
          />
        </>
      )}

      {/* ==================== Update Modal ==================== */}
      {openUpdate && currentTodo && (
        <>
          <div
            className="modal fade show d-block"
            tabIndex={-1}
            role="dialog"
            aria-modal="true"
          >
            <div className="modal-dialog modal-dialog-centered">
              <form
                className="modal-content"
                onSubmit={handleSubmitUpdate(onSubmitUpdate)}
              >
                <div className="modal-header">
                  <h5 className="modal-title">Update Todo</h5>
                  <button
                    type="button"
                    className="btn-close"
                    onClick={() => {
                      setOpenUpdate(false);
                      setCurrentTodo(null);
                    }}
                    disabled={updateLoading}
                  />
                </div>

                <div className="modal-body">
                  <div className="mb-3">
                    <label htmlFor="update-title" className="form-label">
                      Title
                    </label>
                    <input
                      id="update-title"
                      type="text"
                      className={`form-control ${
                        errorsUpdate.title
                          ? "is-invalid"
                          : dirtyUpdate.title
                            ? "is-valid"
                            : ""
                      }`}
                      placeholder="Enter title"
                      disabled={updateLoading}
                      {...registerUpdate("title", {
                        required: "Title is required",
                      })}
                    />
                    {renderFieldError(errorsUpdate.title, dirtyUpdate.title)}
                  </div>

                  <div className="mb-3">
                    <label htmlFor="update-details" className="form-label">
                      Details
                    </label>
                    <textarea
                      id="update-details"
                      className={`form-control ${
                        errorsUpdate.details
                          ? "is-invalid"
                          : dirtyUpdate.details
                            ? "is-valid"
                            : ""
                      }`}
                      placeholder="Enter details"
                      rows={3}
                      disabled={updateLoading}
                      {...registerUpdate("details", {
                        required: "Details is required",
                      })}
                    />
                    {renderFieldError(
                      errorsUpdate.details,
                      dirtyUpdate.details,
                    )}
                  </div>
                </div>

                <div className="modal-footer">
                  <button
                    type="button"
                    className="btn btn-secondary"
                    onClick={() => {
                      setOpenUpdate(false);
                      setCurrentTodo(null);
                    }}
                    disabled={updateLoading}
                  >
                    Close
                  </button>
                  <button
                    type="submit"
                    className="btn btn-primary"
                    disabled={updateLoading}
                  >
                    {updateLoading ? (
                      <>
                        <span
                          className="spinner-border spinner-border-sm me-2"
                          role="status"
                          aria-hidden="true"
                        />
                        Updating...
                      </>
                    ) : (
                      "Update Todo"
                    )}
                  </button>
                </div>
              </form>
            </div>
          </div>
          <div
            className="modal-backdrop fade show"
            onClick={() => {
              if (!updateLoading) {
                setOpenUpdate(false);
                setCurrentTodo(null);
              }
            }}
          />
        </>
      )}
    </div>
  );
}

export default App;
