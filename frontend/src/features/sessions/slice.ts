
import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface sessionState {
  isLoggedIn: boolean;
  userEmail: string | null;
}

const initialState: sessionState = {
  isLoggedIn: false,
  userEmail: null,
};

const sessionSlice = createSlice({
  name: 'session',
  initialState,
  reducers: {
    login(state, action: PayloadAction<{ email: string }>) {
      state.isLoggedIn = true;
      state.userEmail = action.payload.email;
    },
    logout(state) {
      state.isLoggedIn = false;
      state.userEmail = null;
    },
  },
});

export const { login, logout } = sessionSlice.actions;
export default sessionSlice.reducer;
