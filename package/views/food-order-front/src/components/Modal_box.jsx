import React from 'react'
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
const Modal_box = (props) => {
    
  return (
    <div
      className="modal"
      show={props.showModal}
      onHide={props.hideModal}
      style={{ display: "block", position: "initial" }}
    >
      <Modal.Dialog>
        <Modal.Header closeButton>
          <Modal.Title>Dish completion</Modal.Title>
        </Modal.Header>

        <Modal.Body>
          <p>Have you completed the dish ??</p>
        </Modal.Body>

        <Modal.Footer>
          <Button variant="secondary">Close</Button>
          <Button
            onClick={() => onClick(props.order_id, props.food_id)}
            variant="primary"
          >
            Completed
          </Button>
        </Modal.Footer>
      </Modal.Dialog>
    </div>
  );
}

export default Modal_box