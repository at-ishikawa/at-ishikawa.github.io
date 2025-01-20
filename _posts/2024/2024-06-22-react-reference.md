---
title: React reference
date: 2024-06-22
---

The quick reference for a React

## Hooks

A react hook must be called at the top level of React's components.

### useState

### useRef

`useRef` is a useful for a value that’s not needed for rendering, because changing a ref does not trigger a re-render.
This is particularly used to handle DOM elements in React.
See [the officla document](https://react.dev/reference/react/useRef#manipulating-the-dom-with-a-ref) for more details.

```typescript
const Component = () => {
  const fileElement = useRef<HTMLInputElement>(null);
  return (
    <div className={styles.container}>
      <input type="file"
        ref={fileElement}
        onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
          const files = event.target.files;
          // TODO something
        }}
      />
      <button onClick={() => {
        fileElement.current!.click();
      }}>Upload</button>
  )
}
```

### useReducer

[An example code](https://playcode.io/1914386).

```typescript
import {useReducer} from 'react';

interface State {
   count: number
};

type CounterAction =
  | { type: "reset" }
  | { type: "setCount"; value: State["count"] }


const initialState: State = { count: 0 };

function stateReducer(state: State, action: CounterAction): State {
  switch (action.type) {
    case "reset":
      return initialState;
    case "setCount":
      return { ...state, count: action.value };
    default:
      throw new Error("Unknown action");
  }
}

export default function App() {
  const [state, dispatch] = useReducer<State>(stateReducer, initialState);

  const increment = () => dispatch({ type: "setCount", value: state.count+1 });
  const reset = () => dispatch({ type: "reset" });

  return (
    <div>
      <h1>Welcome to my counter</h1>

      <p>Count: {state.count}</p>
      <button onClick={increment}>Add</button>
      <button onClick={reset}>Reset</button>
    </div>
  );
}
```

### useContext

[Example](https://playcode.io/1914368)

The example code for a theme context.

```typescript
import { createContext, useContext } from 'react';

type Theme = "light" | "dark" | "system";

// Create a context
const ThemeContext = createContext<Theme>("system");

const useGetTheme = () => useContext(ThemeContext);

export default function App() {
  return (
    <ThemeContext.Provider value={"system"}>
      <MyComponent />
    </ThemeContext.Provider>
  )
}

function MyComponent() {
  const theme = useContext(ThemeContext);
  // const theme = useGetTheme();

  return (
    <div className="App">
      <h1>Current theme: {theme}</h1>
    </div>
  )
}
```

### useCallback

Cache a function result between re-renders.
See [this reference](https://react.dev/reference/react/useCallback).
This is especially necessary with `memo` to pass the same callback function without changing a reference.


### useEffect

See [the official doc](https://react.dev/reference/react/useEffect) for the details.

```typescript
interface CleanupFunction {
    () => void
}
interface SetupFunction {
    () => CleanupFunction
}
function setUp(): SetUpFunction {
    //
    return () => {
        // cleanup function
    }
}
useEffect(setUp, [dependency1, dependency2]];
```

This `useEffect` can run when it's mounted or updated.
If we want to run it only when the component is updated, then we can use `useRef` to check if it's mounted or not.
See [this Stackoverflow answer](https://stackoverflow.com/a/55075818) for more details.


## Higher Order Components

It's common to use a higher order function to add an extra behavior to a component in React.

For example, to check if a user is authenticated or not, and do something else like redirect a page, then one higher order function may be used.
For the typescript, [this article](https://medium.com/@jrwebdev/react-higher-order-component-patterns-in-typescript-42278f7590fb) can be useful.

```typescript
interface withAuthenticatedProps {
  authenticatedUser: {
    name: string;
  }
}
const withAuthenticatedUser = <Props extends object>(WrappedComponent: React.ComponentType<Props & withAuthenticatedProps>) => {
  return (props: Props) => {
    // TODO replace with a context
    const [authenticatedUser, setAuthenticatedUser] = React.useState({
      name: "User"
    })
    if (authenticatedUser == null) {
      // TODO: Replace a logic with more appropriate form like a redirect
      return <>Not authenticated</>
    }
    return <WrappedComponent {...props} authenticatedUser={authenticatedUser} />
  }
}

export default withAuthenticatedUser(Component);
```


## TypeScript

- Use either `React.ReactNode` for the children of a react component.
- Use `React.ComponentType<P>` (`ComponentClass<P> | FunctionComponent<P>`) for a higher order component.


### DOM elements and events

- input element: `HTMLInputElement`
    - event of the onChange: `React.ChangeEvent<HTMLInputElement>`
- form element: `HTMLFormElement`
    - event of the onSubmit: `React.FormEvent<HTMLFormElement>`
