import React from 'react'
import './LoginSignup.css'
import user_icon from '../Assets/person.png'
import email_icon from '../Assets/email.png'
import password_icon from '../Assets/password.png'
const LoginSignup = () => {
  return (
    <div className='Container'>
      <div className='header'>
        <div className='text'>Signup</div>
        <div className='underline'></div>
      </div>
      <div className='inputs'>
        <div className='input'>
            <img src={user_icon} alt='Username' />
            <input type='text' placeholder='Username'/>
        </div>
        <div className='input'>
            <img src={email_icon} alt='Email' />
            <input type='email' placeholder='Email'/>
        </div>
        <div className='input'>
            <img src={password_icon} alt='Password' />
            <input type='password' placeholder='Password'/>
        </div>
      </div>
      <div className='ForgotPassword'> Lost password? <span>Reset</span></div>
      <div className='SubmitContainer'>
        <div className='signup'>Signup</div>
        <div className='login'>Login</div>
      </div>
    </div>
  )
}

export default LoginSignup
