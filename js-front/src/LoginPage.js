import { Col, Button, Form, Container, Row } from "react-bootstrap";
import { useHook } from './util';
import { isLoggedIn } from './auth';
import { Redirect } from "react-router-dom";

export default function LoginPage(props) {
    const [email, , emailProps] = useHook('admin@admin.pl');
    const [password, , passwordProps] = useHook('admin');

    function handleSubmit(event) {
        event.preventDefault();
        fetch('http://localhost:8090/login', {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email, password
            }),
            credentials: "include"
        }).then(res => {
            if (res.ok) {
                props.history.push('/note');
            } else {
                alert("Wrong credentials");
            }
        });
    }

    if (isLoggedIn()) {
        return (
            <Redirect to="/note" />
        );
    }

    return (
        <Container className="text-light">
            <h1>Login</h1>
            <Form onSubmit={handleSubmit}>
                <Row>
                    <Col sm="12" md="6">
                        <Form.Group>
                            <Form.Label>Email</Form.Label>
                            <Form.Control {...emailProps} type="email" />
                        </Form.Group>
                    </Col>
                    <Col sm="12" md="6">
                        <Form.Group>
                            <Form.Label>Password</Form.Label>
                            <Form.Control {...passwordProps} type="password" />
                        </Form.Group>
                    </Col>
                </Row>
                <Button type="submit">Submit</Button>
            </Form>
        </Container>
    )
}