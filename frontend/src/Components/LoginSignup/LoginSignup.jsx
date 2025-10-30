import React from 'react'
import './LoginSignup.css'
import user_icon from '../Assets/person.png'
import email_icon from '../Assets/email.png'
import password_icon from '../Assets/password.png'
import { signup, login } from '../../api/auth'

const LoginSignup = () => {
  const [action, setAction] = React.useState('Sign up')
  const [form, setForm] = React.useState({ username: '', email: '', password: '' })
  const [isSubmitting, setIsSubmitting] = React.useState(false)
  const [message, setMessage] = React.useState('')

  const onChange = (e) => {
    const { name, value } = e.target
    setForm((f) => ({ ...f, [name]: value }))
  }

  const handleSignup = async () => {
    setMessage('')
    setIsSubmitting(true)
    try {
      const res = await signup(form)
      setMessage(res?.message || 'Signup successful')
    } catch (err) {
      setMessage(err?.response?.data?.error || err.message || 'Signup failed')
    } finally {
      setIsSubmitting(false)
    }
  }

  const handleLogin = async () => {
    setMessage('')
    setIsSubmitting(true)
    try {
      const res = await login({ email: form.email, password: form.password })
      setMessage(res?.message || 'Login successful')
    } catch (err) {
      setMessage(err?.response?.data?.error || err.message || 'Login failed')
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <div className='Container'>
      <div className='header'>
        <div className='text'>{action}</div>
        <div className='underline'></div>
      </div>

      <div className='inputs'>
        {action === 'Login' ? null : (
          <div className='input'>
            <img src={user_icon} alt='Username' />
            <input
              name='username'
              type='text'
              placeholder='Username'
              value={form.username}
              onChange={onChange}
            />
          </div>
        )}

        <div className='input'>
          <img src={email_icon} alt='Email' />
          <input
            name='email'
            type='email'
            placeholder='Email'
            value={form.email}
            onChange={onChange}
          />
        </div>

        <div className='input'>
          <img src={password_icon} alt='Password' />
          <input
            name='password'
            type='password'
            placeholder='Password'
            value={form.password}
            onChange={onChange}
          />
        </div>
      </div>

      {action === 'Sign up' ? null : (
        <div className='ForgotPassword'>
          Lost password? <span>Reset</span>
        </div>
      )}

      <div style={{ textAlign: 'center', color: '#3c009d', minHeight: 24 }}>{message}</div>

      <div className='SubmitContainer'>
        <div
          className={`Submit ${isSubmitting ? 'gray' : ''}`}
          onClick={() => {
            if (isSubmitting) return
            if (action !== 'Sign up') setAction('Sign up')
            else handleSignup()
          }}
        >
          Sign up
        </div>

        <div
          className={`Submit ${isSubmitting ? 'gray' : ''}`}
          onClick={() => {
            if (isSubmitting) return
            if (action !== 'Login') setAction('Login')
            else handleLogin()
          }}
        >
          Login
        </div>
      </div>
    </div>
  )
}

export default LoginSignup
