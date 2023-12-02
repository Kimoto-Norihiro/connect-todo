import { createPromiseClient } from "@bufbuild/connect";
import { TODOService } from '../api/todoservice/v1/todo_connectweb';
import { createConnectTransport } from '@bufbuild/connect-web';
import { TODO } from '@/api/todoservice/v1/todo_pb'

const transport = createConnectTransport({
	baseUrl: "http://localhost:8080",
});
const client = createPromiseClient(TODOService, transport);

export const createTodo = async () => {
	const res = await client.createTODO({
		title: "new todo",
	} as TODO)
	console.log(res)
};

export const listTodos = async (setTodos: React.Dispatch<React.SetStateAction<TODO[]>>) => {
	try {
		const res = await client.listTODOs({})
		setTodos(res.todos)
	} catch (e) {
		console.log(e)
	}
}

export const deleteTodo = async (id: number) => {
	const res = await client.deleteTODO({
		id,
	})
	console.log(res)
}

export const updateTodo = async (todo: TODO) => {
	const res = await client.updateTODO({
		todo: todo,
	})
	console.log(res)
}