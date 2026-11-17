import API from "./index";

import type {
  TodoInterface,
  CreateTodoInterface,
  UpdateTodoInterface,
} from "../types/todo";

const baseURL = "/todos";

export async function listTodos(): Promise<TodoInterface[]> {
  const response = await API.get<TodoInterface[]>(baseURL);

  return response.data;
}

export async function getTodo(todoID: string): Promise<TodoInterface> {
  const response = await API.get<TodoInterface>(`${baseURL}/${todoID}`);

  return response.data;
}

export async function createTodo(
  todoData: CreateTodoInterface,
): Promise<TodoInterface> {
  const response = await API.post<TodoInterface>(baseURL, todoData);

  return response.data;
}

export async function updateTodo(
  todoID: string,
  todoData: UpdateTodoInterface,
): Promise<TodoInterface> {
  const response = await API.patch<TodoInterface>(
    `${baseURL}/${todoID}`,
    todoData,
  );

  return response.data;
}

export async function deleteTodo(todoID: string): Promise<void> {
  await API.delete(`${baseURL}/${todoID}`);
}
