import React from 'react'
import { Logo } from './Logo'

export const Header: React.FC = () => {
  return (
    <header>
      <div className='flex items-center space-x-2 md:space-x-10'>
        <Logo />
      </div>
    </header>
  )
}
