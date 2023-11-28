import React, { useState } from "react";
import { useCommonModal } from "../../contexts/modal-context";
import { DeleteModal } from "./DeleteModal";
import { TODO } from "@/api/todoservice/v1/todo_pb";
import { deleteTodo, listTodos, updateTodo } from "@/handlers/todo_handlers";

type Props = {
  todo: TODO
  setTodos: React.Dispatch<React.SetStateAction<TODO[]>>
}

const TodoCard = ({ todo, setTodos }: Props) => {
  const { showModal } = useCommonModal();
  const [done, setDone] = useState(false)
  const [isEdit, setIsEdit] = useState(false)
  const [title, setTitle] = useState(todo.title)

  const onEnterDown = async (e: React.KeyboardEvent) => {
		if (e.key == "Enter") {
      await updateTodo({
        id: todo.id,
        title,
      } as TODO)
      setIsEdit(false)
      listTodos(setTodos)
    }
  }

  return (
    <div className="flex rounded-lg p-4 py-2 items-center justify-between border-2 border-gray-500 m-1">
      <div className="flex items-center w-1/2">
        <div 
          className={`h-3 w-3 rounded-lg ${done ? 'bg-green-500' : 'border-2 border-black'}`}
          onClick={() => setDone(!done)}
        ></div>
        {
          isEdit ? (
            <input 
              className='rounded-sm px-1 text-lg font-bold ml-3' 
              value={title} 
              onChange={(e) => setTitle(e.target.value)}
              onKeyDown={onEnterDown}
              onBlur={() => {
                setIsEdit(false)
              }}
            />
          ) : (
            <h2 
              className={`text-lg font-bold ml-4 ${done && 'line-through'}`}
              onClick={() => setIsEdit(true)}
            >
              {todo.title}
            </h2>
          )
        }
      </div>
      <div className="flex">
        {
          done && (
            <div 
              className="text-red-500 hover:text-red-700 ml-4"
              onClick={() => 
                showModal(
                  <DeleteModal 
                    message="are you really delete it ?"
                    deleteAction={() => deleteTodo(todo.id)}
                  />
                )
              }
            >
              delete
            </div>
          )
        }
      </div>
    </div>
  );
};

export default TodoCard;