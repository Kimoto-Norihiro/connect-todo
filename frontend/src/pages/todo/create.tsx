import { NextPage } from 'next'
import React, { useEffect, useState } from 'react'
import { InputWithError } from '../../components/parts/InputWithError'
import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { useRouter } from 'next/router'
import axios from 'axios'
import * as yup from 'yup';
import { TODO } from '@/api/todoservice/v1/todo_pb'

type TodoCreateRequest = Pick<TODO, 'title'>

const todoSchema = yup.object().shape({
  title: yup.string().required('required input'),
})

const CreateTodo: NextPage = () => {
  const [err, setErr] = useState('')
  const router = useRouter()

  const { register, handleSubmit, formState: { errors } } = useForm<TodoCreateRequest>({
    resolver: yupResolver(todoSchema),
  });

  const createTodo = async (todo: TodoCreateRequest) => {
    
  }

  useEffect(() => {
    console.log('authToken', localStorage.getItem('authToken'))
  }, [])

  const onSubmit = async () => {
    handleSubmit(async (data) => {
      const response = await createTodo({
        title: data.title,
      })
      console.log(response)
      router.back()
    }, (err) => {
      console.log(err)
      console.log('error')
    })()
  }

  // useEffect(() => {
  //   if (!authToken) router.push('/login')
  // }, [authToken])

  return (
    <div className='flex flex-col bg-white justify-center align-center p-10'>
      <div className='flex justify-center text-3xl'>
        <h1>Create Todo</h1>
      </div>
      <div className='flex p-5 justify-center bg-white'>
        <form
          className='w-96 p-3' 
          onSubmit={(e) => {
            e.preventDefault()
            onSubmit()
          }}
        >
          <InputWithError 
            name='title'
            register={register}
            errors={errors}
          />
          {
            err ? <div className='text-red-800 text-sm'>{err}</div> : <div className='h-5'></div>
          }
          <div className='flex justify-center mt-5'>
            <div className='bg-gray-200 py-1 px-3 rounded-lg'>
              <button type='submit'>submit</button>
            </div>
          </div>
        </form>
      </div>
    </div>
  )
}


export default CreateTodo