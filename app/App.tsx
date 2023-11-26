import React, { useEffect, useState } from 'react';

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

interface Todo {
    id: number;
    title: string;
    isDone: boolean;
}

const Todo: React.FC<Todo> = ({ title, isDone, id}) => {
    return (
        <li>
            <input type="checkbox" checked={isDone} id={id.toString()}></input>
            <label htmlFor={id.toString()}>{title}</label>
        </li>
    );
};

const App: React.FC = () => {
    const [text, setText] = useState('');
    const [todos, setTodos] = useState<Todo[]>([]);

    const Todos = todos.map((todo) => (<Todo {...todo} />));

    useEffect(() => {
        console.log("hello")
        fetch('http://localhost:3000/api/todos')
            .then((res) => res.json())
            .then((data) => {
                console.log(data);
                setTodos(data);
            });
    }, []);
    const logIn = (formData: FormData) => {
        const email = formData.get('email');
        const password = formData.get('password');
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ password: password, email: email }),
        };
        fetch('http://localhost:3000/api/login', requestOptions).then(
            (response) => getSecret()
        );
    };

    const getSecret = () => {
        const requestOptions = {
            method: 'GET',
        };
        fetch('http://localhost:3000/api/secret', requestOptions)
            .then((req) => req.text())
            .then((text) => setText(text));
    };

    return (
        <div>
            {/* @ts-ignore */}
            <form action={logIn}>
                <Input name="email" label="E-mail" type="email" />
                <Input name="password" label="Password" type="password" />
                <button>Log In</button>
                {text}
                <div>
                    {Todos}
                </div>
            </form>
        </div>
    );
};

export default App;
