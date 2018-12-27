import {ActionType, createSessionAction} from "./login.actions";
import {initialLoginState} from "./login.state";
import * as api from '../api'
import {createMemoryHistory} from "history";

jest.mock('../api', () => ({
    createSession: jest.fn()
}));

describe('login actions', () => {
    const history = createMemoryHistory();

    it('should dispatch login success', async () => {
        expect.assertions(1);

        const dispatch = jest.fn();
        const state = {
            ...initialLoginState,
            username: "a@b.com",
            password: "Password123"
        };
        (api.createSession as jest.Mock).mockReturnValue(
            Promise.resolve({
                isSuccess: jest.fn(() => true),
            }));

        await createSessionAction(dispatch, state, history)();

        expect(dispatch.mock.calls).toEqual([
            [{type: ActionType.SubmitCredentials}],
            [{type: ActionType.LoginSuccess}],
        ]);
    });
});
