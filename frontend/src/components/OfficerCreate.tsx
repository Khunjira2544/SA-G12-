import React from "react";
import { Link as RouterLink } from "react-router-dom";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import { OfficersInterface } from "../models/IOfficer";
import { CreateOfficer } from "../services/HttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function OfficerCreate() {
    const [officer, setOfficer] = React.useState<Partial<OfficersInterface>>({});
    const [success, setSuccess] = React.useState(false);
    const [error, setError] = React.useState(false);

    const handleClose = (
        event?: React.SyntheticEvent | Event,
        reason?: string
    ) => {
        if (reason === "clickaway") {
            return;
        }
        setSuccess(false);
        setError(false);
    };

    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof OfficerCreate;
        const { value } = event.target;
        setOfficer({ ...officer, [id]: value });
    };

    async function submit() {
        let res = await CreateOfficer(officer);
        if (res) {
            setSuccess(true);
        } else {
            setError(true);
        }
    }

    return (
        <Container maxWidth="md">
            <Snackbar open={success} autoHideDuration={3000} onClose={handleClose} anchorOrigin={{ vertical: "top", horizontal: "center" }}>
                <Alert onClose={handleClose} severity="success">
                    ??????????????????????????????????????????????????????
                </Alert>
            </Snackbar>
            <Snackbar open={error} autoHideDuration={6000} onClose={handleClose} anchorOrigin={{ vertical: "top", horizontal: "center" }}>
                <Alert onClose={handleClose} severity="error">
                    ???????????????????????????????????????????????????????????????
                </Alert>
            </Snackbar>
            <Paper>
                <Box display="flex" sx={{ marginTop: 2, }}>
                    <Box sx={{ paddingX: 2, paddingY: 1 }}>
                        <Typography component="h2" variant="h6" color="primary" gutterBottom>
                            ???????????????????????????????????????????????????
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}>
                    <Grid item xs={6}>
                        <p>???????????? - ????????????</p>
                        <FormControl fullWidth variant="outlined">
                            <TextField id="Name" variant="outlined" type="string" size="medium" placeholder="?????????????????????????????????????????????????????????" value={officer.Name || ""} onChange={handleInputChange}/>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>???????????????</p>
                            <TextField id="Email" variant="outlined"  type="string" size="medium" placeholder="????????????????????????????????????????????????????????????" value={officer.Email || ""} onChange={handleInputChange}/>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>????????????????????????</p>
                            <TextField id="Password" variant="outlined" type="password" size="medium" placeholder="???????????????????????????????????????????????????" value={officer.Password || ""} onChange={handleInputChange}/>
                        </FormControl>
                    </Grid>
                    <Grid item xs={12}>
                        <Button component={RouterLink} to="/Officers" variant="contained" color="inherit">
                            ????????????
                        </Button>
                        <Button style={{ float: "right" }} onClick={submit} variant="contained" color="primary">
                            ??????????????????
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
    );
}

export default OfficerCreate;