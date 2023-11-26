import React, {useState} from 'react';

interface InputProps {
    name: string;
    label: string;
    type?: string;
}

const Input: React.FC<InputProps> = ({ name, label, type }) => (
    <div>
        <label htmlFor={name}>{label}</label>
        <input name={name} id={name} type={type} />
    </div>
);

const App: React.FC = () => {
    const [text, setText] = useState("")
    function logIn(formData: FormData) {
        const email = formData.get('email');
        const password = formData.get('password');
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ password: password, email: email }),
        };
        fetch('http://localhost:3000/login', requestOptions)
            .then((response) => setText(response.status.toString()))
    }

    return (
        // @ts-ignore
        <form action={logIn}>
            <Input name="email" label="E-mail" type="email" />
            <Input name="password" label="Password" type="password" />
            <button>Log In</button>
            {text}
        </form>
    );
};

export default App;
