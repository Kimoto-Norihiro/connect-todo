import { GetServerSideProps, NextPage } from 'next'
import React, { useEffect, useState } from 'react'
import { User } from '../types/types'
import { useRouter } from 'next/router'
import { authAxios } from '@/types/api/axios'
import TodoCard from '../components/parts/TodoCard'
import { useCommonModal } from '../contexts/modal-context';
import { DeleteModal } from '../components/parts/DeleteModal'

const MyPage: NextPage = () => {
  const router = useRouter()
  const { showModal } = useCommonModal()
  const [user, setUser] = useState<User>()

  const getUser = async () => {
    
  }

  const deleteUser = async () => {
    
  }

  return (
    <div className='flex flex-col bg-white justify-center align-center'>
      <div className='mx-auto container'>
        <p className='text-3xl my-2'>Todo List</p>
        <div className='border-2 rounded-lg border-gray-500 p-6'>
          {
            user?.todos && user.todos.length !== 0 ? (
              user.todos.map((todo, index) => {                
                return (    
                  <TodoCard key={index} todo={todo}/>
                )
              })
            ) : (
              <p>no todo</p>
            )
          }
          <div onClick={() => router.push('/todo/create')}>
            <p className='text-blue-500 hover:text-blue-700'>+ add todo</p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default MyPage