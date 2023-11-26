import { NextPage, GetStaticPaths, GetStaticProps } from 'next'
import React, { useEffect, useState } from 'react'
import { InputWithError } from '../../components/parts/InputWithError'
import { useForm } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import * as yup from 'yup';
import axios from 'axios'
import { useRouter } from 'next/router'
import { TODO } from '@/api/todoservice/v1/todo_pb'

type TodoUpdateRequest = Pick<TODO, 'title'>

type Props = {
  id: string
}

const todoSchema = yup.object().shape({
  title: yup.string().required('required input'),
})

const EditTodo: NextPage<Props> = ({ id }) => {
  const router = useRouter()
  const [err, setErr] = useState('')
  const [todo, setTodo] = useState<TodoUpdateRequest>()
  const { register, handleSubmit, formState: { errors }, setValue } = useForm<TodoUpdateRequest>({
    resolver: yupResolver(todoSchema),
  });

  const getTags = async () => {
    
  }

  const getTodo = async () => {
   
  }

  const updateTodo = async (todo: TodoUpdateRequest) => {
    
  }

  const onSubmit = async () => {
    handleSubmit(async (data) => {
      const response = await updateTodo({
        title: data.title,
      })
      console.log(response)
      router.back()
    }, (err) => {
      console.log(err)
      console.log('error')
    })()
  }

  useEffect(() => {
    if (todo) {
      setValue('title', todo?.title || '')
    }
  },[todo])

  return (
    <div className='flex flex-col bg-white justify-center align-center p-10'>
      <div className='flex justify-center text-3xl'>
        <h1>Edit Todo</h1>
      </div>
      <div className='flex p-5 justify-center bg-white'>
        <form
          className='w-96 p-3' 
          onSubmit={(e) => {
          e.preventDefault()
          onSubmit()
        }}>
          <InputWithError 
            name='title'
            register={register}
            errors={errors}
          />
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

type PathParams = {
  id: string
}

export const getStaticPaths: GetStaticPaths<PathParams> = async () => {
  return {
    paths: [],
    fallback: true,
  };
};

export const getStaticProps: GetStaticProps = async ({params}) => {
  const { id } = params as PathParams
  return {
    props: {
      id,
    },
  }
}

export default EditTodo