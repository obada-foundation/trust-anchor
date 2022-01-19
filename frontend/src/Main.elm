module Main exposing (main)

import Browser
import Element exposing (..)

type alias Flags =
    {}

type Msg
    = NoOp

type alias Model =
    ()

main : Program Flags Model Msg
main =
    Browser.document
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }

initialModel : Model
initialModel =
    ()

init : Flags -> ( Model, Cmd Msg )
init _ =
    ( initialModel, Cmd.none )

subscriptions : Model -> Sub Msg
subscriptions _ =
    Sub.none

view : Model -> Browser.Document Msg
view _ =
    { title = "TrustAnchor | Main"
    , body =
        [ layout
            []
            (text "TrustAnchor UI")
        ]
    }

update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    -- case Debug.log "msg" msg of
    case msg of
        NoOp ->
            ( model, Cmd.none )
