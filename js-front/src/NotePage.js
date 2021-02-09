import { useEffect, useState } from "react";
import { Container, Form, Button, Row } from "react-bootstrap";
import { logout } from "./auth";
import { myFetch, myFetchPost, useInput } from "./util";

export default function NotePage({ history }) {
    const [note, setNote, noteProps] = useInput('');
    const [prevNote, setPrevNote] = useState('');

    useEffect(() => {
        myFetch("note").then(val => {
            setNote(val);
            setPrevNote(val);
        });
    }, []);

    function handleSubmit(e) {
        e.preventDefault();
        myFetchPost("note", { note }).then(() => setPrevNote(note));
    }

    const button = (
        note === prevNote
            ? <Button variant="success" className="mt-2 mr-2 disabled" style={{ width: '7rem' }}>Saved</Button>
            : <Button className="mt-2 mr-2" style={{ width: '7rem' }} type="submit">Save</Button>
    );

    return (
        <div className="NotePage">
            <Container>
                <Row className="justify-content-center">
                    <Form onSubmit={handleSubmit} style={{ maxWidth: '768px', width: '100%' }}>
                        <Form.Control as="textarea" {...noteProps} />
                        {button}
                        <Button className="mt-2 mr-2" variant="warning" onClick={() => logout(history)}>Log Out</Button>
                    </Form>
                </Row>
            </Container>
        </div>
    )
}