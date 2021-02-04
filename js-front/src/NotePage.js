import { useEffect } from "react";
import { Container, Form, Button } from "react-bootstrap";
import { myFetch, useHook } from "./util";

export default function NotePage() {
    const [note, setNote, noteProps] = useHook('');

    useEffect(() => {
        myFetch("http://localhost:8090/note", { credentials: 'include' }).then(val => {
            setNote(val);
        })
    }, []);

    function handleSubmit(e) {
        e.preventDefault();
        fetch("http://localhost:8090/note", {
            method: "POST",
            credentials: 'include',
            body: JSON.stringify({ note }),
            headers: {
                "Content-Type": "application/json"
            }
        }).then(res => {
            if (res.ok) {
                alert("Git");
            } else {
                alert(res.status);
            }
        })
    }

    return (
        <div className="NotePage">
            <Container>
                <Form onSubmit={handleSubmit}>
                    <Form.Control as="textarea" {...noteProps} />
                    <Button type="submit">Submit</Button>
                </Form>
                <Button onClick={(e) => {

                }}>Log Out</Button>
            </Container>
        </div>
    )
}