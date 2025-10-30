import api from './client'

export async function signup({ username, email, password }) {
  const { data } = await api.post('/signup', { username, email, password })
  return data
}

export async function login({ email, password }) {
  const { data } = await api.post('/login', { email, password })
  return data
}