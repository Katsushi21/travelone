import { Button, Stack } from '@chakra-ui/react';

export const Submit = () => {
  return;
  <Stack direction="row" spacing={4} align="center">
    <Button
      isLoading
      loadingText="Loading"
      colorScheme="teal"
      variant="outline"
      spinnerPlacement="start"
    >
      Submit
    </Button>
  </Stack>;
};
