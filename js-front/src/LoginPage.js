import { Col, Button, Form, Container, Row } from "react-bootstrap";
import { myFetchPost, useInput } from './util';
import { isLoggedIn } from './auth';
import { Redirect } from "react-router-dom";

export default function LoginPage(props) {
    const [email, , emailProps] = useInput('admin@admin.pl');
    const [password, , passwordProps] = useInput('admin');

    function handleSubmit(event) {
        event.preventDefault();
        myFetchPost('login', { email, password })
            .then(() => props.history.push('/note'))
            .catch(alert);
    }

    if (isLoggedIn()) {
        return (
            <Redirect to="/note" />
        );
    }

    return (
        <Container className="text-light">
            <Row>
                <Col className="text-center">
                    <h1 className="mb-3" style={{ fontSize: 'calc(1.525rem + 3.3vw)', lineHeight: "1" }}>Note</h1>
                    <p className="lead">Use any email you'd like - if it hasn't been used before a new note will be created anyways</p>
                    <p className="lead">Password - admin</p>
                </Col>
            </Row>
            <Row className="justify-content-center">
                <Form onSubmit={handleSubmit} style={{ maxWidth: '768px', width: '100%' }}>
                    <Form.Group>
                        <Form.Label>Email</Form.Label>
                        <Form.Control {...emailProps} type="email" />
                    </Form.Group>
                    <Form.Group>
                        <Form.Label>Password</Form.Label>
                        <Form.Control {...passwordProps} type="password" />
                    </Form.Group>
                    <Button type="submit">Submit</Button>
                </Form>
            </Row>
        </Container>
    )
}