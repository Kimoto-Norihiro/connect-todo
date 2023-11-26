import { GetServerSideProps, NextPage } from 'next'
import React, { use, useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import TodoCard from '../components/parts/TodoCard'
import { useCommonModal } from '../contexts/modal-context';
import { createPromiseClient } from "@bufbuild/connect";
import { TODOService } from '../api/todoservice/v1/todo_connectweb';
import { createConnectTransport } from '@bufbuild/connect-web';
import { TODO } from '@/api/todoservice/v1/todo_pb'

const MyPage: NextPage = () => {
  const router = useRouter()
  const { showModal } = useCommonModal()
  const [todos, setTodos] = useState<TODO[]>([])

  const transport = createConnectTransport({
    baseUrl: "http://localhost:8080",
  });
  const client = createPromiseClient(TODOService, transport);

  const createTodo = async () => {
    const res = await client.createTODO({
      title: "test",
    })
    console.log(res)
  };
  
  const listTodos = async () => {
    try {
      const res = await client.listTODOs({})
      console.log(res)
    } catch (e) {
      console.log(e)
    }
  }

  const deleteTodo = async (id: number) => {
    const res = await client.deleteTODO({
      id: 1,
    })
    console.log(res)
  }

  const updateTodo = async () => {
    const res = await client.updateTODO({
      todo: {
        id: 1,
        title: "test",
      }
    })
    console.log(res)
  }

  useEffect(() => {
    listTodos()
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
                  <TodoCard key={index} todo={todo}/>
                )
              })
            ) : (
              <p>no todo</p>
            )
          }
          <div onClick={createTodo}>
            <p className='text-blue-500 hover:text-blue-700'>+ add todo</p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default MyPage