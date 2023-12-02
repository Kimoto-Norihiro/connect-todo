import { GetServerSideProps, NextPage } from 'next'
import React, { use, useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import TodoCard from '../components/parts/TodoCard'
import { useCommonModal } from '../contexts/modal-context';
import { createPromiseClient } from "@bufbuild/connect";
import { TODOService } from '../api/todoservice/v1/todo_connectweb';
import { createConnectTransport } from '@bufbuild/connect-web';
import { TODO } from '@/api/todoservice/v1/todo_pb'
import { listTodos, createTodo } from '@/handlers/todo_handlers';

const MyPage: NextPage = () => {
  const [todos, setTodos] = useState<TODO[]>([])

  useEffect(() => {
    listTodos(setTodos)
  }, [])

  return (
    <div className='flex flex-col bg-white justify-center align-center'>
      <div className='mx-auto container'>
        <p className='text-3xl my-2'>Todo List</p>
        <div className='border-2 rounded-lg border-gray-500 p-6'>
          {
            todos && todos.length !== 0 ? (
              todos.map((todo, index) => {                
                return (    
                  <TodoCard key={index} todo={todo} setTodos={setTodos}/>
                )
              })
            ) : (
              <p>no todo</p>
            )
          }
          <div onClick={async () => {
            await createTodo()
            listTodos(setTodos)
          }}>
            <p className='text-blue-500 hover:text-blue-700'>+ add todo</p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default MyPage