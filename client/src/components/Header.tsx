import React from 'react'
import { Logo } from './Logo'
import { BellIcon, SearchIcon } from '@heroicons/react/solid'

export const Header: React.FC = () => {
  return (
    <header>
      <div className='flex items-center space-x-2 md:space-x-10'>
        <Logo />
        <ul className='hidden space-x-4 md:flex'>
          <li className='headerLink'>a</li>
          <li className='headerLink'>b</li>
          <li className='headerLink'>c</li>
          <li className='headerLink'>d</li>
          <li className='headerLink'>e</li>
        </ul>
      </div>
      <div className='flex item-center space-x-4'>
        <SearchIcon className='hidden h-6 w-6 sm:inline' />
        <p className='hidden lg:inline'>Kids</p>
        <BellIcon className='h-6 w-6' />
      </div>
    </header>
  )
}
