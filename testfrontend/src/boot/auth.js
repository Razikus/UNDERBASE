import { AuthClient } from '@supabase/auth-js'
import { boot } from 'quasar/wrappers'


const GOTRUE_URL = process.env.AUTHORIZER

const authClient = new AuthClient({ url: GOTRUE_URL })
export default boot(async (/* { app, router, ... } */) => {
})

export { authClient }
