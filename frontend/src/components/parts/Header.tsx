import React from 'react'
import { useRouter } from 'next/router'

export const Header = () => {
  const router = useRouter()
  return (
    <div className='bg-black flex py-2'>
      <div className='mx-auto container flex items-center justify-between'>
        <p className='text-white font-bold text-3xl'>Todo app</p>
      </div>
    </div>
  )
}