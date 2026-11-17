export interface TodoInterface {
  id: string;
  title: string;
  details: string;
  isDone: boolean;
}

export interface CreateTodoInterface {
  title: string;
  details: string;
}

export interface UpdateTodoInterface {
  title?: string;
  details?: string;
  // isDone?: boolean;
}
