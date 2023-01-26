function A() {
    $1
}
function IsIssuable() {
    echo ${FUNCNAME[0]}
    echo ${FUNCNAME:2}
}

function main() {

    A IsIssuable
}

function setOrg() {

    Org_Num=1

    ############ Peer Vasic Setting ############
    export CORE_PEER_TLS_ENABLED=true
    export CORE_PEER_ID="$PEERID.org${Org_Num}.example.com"
    export CORE_PEER_ENDORSER_ENABLED=true
    export CORE_PEER_ADDRESS="$PEERID.org${Org_Num}.example.com:7050"
    export CORE_PEER_LOCALMSPID="Org${Org_Num}MSP"
    export CORE_PEER_TLS_ROOTCERT_FILE="${TEST_NETWORK_HOME}/organizations/peerOrganizations/org${Org_Num}.example.com/peers/peer0.org${Org_Num}.example.com/tls/ca.crt"
    export CORE_PEER_MSPCONFIGPATH="${TEST_NETWORK_HOME}/organizations/peerOrganizations/org${Org_Num}.example.com/users/Admin@org${Org_Num}.example.com/msp"
    export ORDERER_CA="${TEST_NETWORK_HOME}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/ca.crt"
}

main
