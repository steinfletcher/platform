import {shallow} from 'enzyme';

import {LoginForm} from "./Login.component";
import * as React from "react";
import {noop} from "../shared/lang";

const noopActions = {
    updateUsername: noop,
    updatePassword: noop,
    onCreateSession: noop,
};

describe('LoginForm', () => {
    it('should render', () => {
        const wrapper = shallow(
            <LoginForm
                disabled={false}
                username={"a@b.com"}
                password={"Password123"}
                loginPending={false}
                {...noopActions}
            />);

        expect(wrapper).toMatchSnapshot();
    });

    it('should disable submit when login pending', () => {
        const wrapper = shallow(
            <LoginForm
                disabled={true}
                username={"a@b.com"}
                password={"Password123"}
                loginPending={true}
                {...noopActions}
            />);

        const submitButton = wrapper.find('[data-test-id="submitBtn"]');

        expect(submitButton.prop("disabled")).toBeTruthy();
        expect(submitButton.text()).toContain('Loading...');
    });

    it('should display an error when error message is present', () => {
        const wrapper = shallow(
            <LoginForm
                disabled={false}
                username={"a@b.com"}
                password={"Password123"}
                loginPending={false}
                errorMessage={"wrong password"}
                {...noopActions}
            />);

        const errorMessage = wrapper.find('[data-test-id="incorrectCredentials"]');

        expect(errorMessage.text()).toEqual("wrong password");
    });
});
