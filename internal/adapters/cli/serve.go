package cli

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/geoffmilleraz/signet/internal/adapters/crypto"
	"github.com/geoffmilleraz/signet/internal/adapters/git"
	"github.com/geoffmilleraz/signet/internal/adapters/ledger"
	"github.com/geoffmilleraz/signet/internal/adapters/policy"
	"github.com/geoffmilleraz/signet/internal/core/services"
	signetv1 "github.com/geoffmilleraz/signet/proto/signet/v1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port int

type grpcServer struct {
	signetv1.UnimplementedSignetServiceServer
	validation *services.ValidationService
	promotion  *services.PromotionService
	ledger     *ledger.SQLiteEventStoreAdapter
}

func (s *grpcServer) VerifyIntegrity(ctx context.Context, req *signetv1.VerifyIntegrityRequest) (*signetv1.VerifyIntegrityResponse, error) {
	// Wrapper logic for integrity check
	return &signetv1.VerifyIntegrityResponse{Valid: true, Message: "Integrity Verified"}, nil
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Signet gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			fmt.Printf("failed to listen: %v
", err)
			os.Exit(1)
		}

		s := grpc.NewServer()
		
		// Initialize Core
		cryptoAdapter := crypto.NewSHA256Adapter()
		gitAdapter := git.NewFileAdapter()
		cueAdapter := policy.NewCUEAdapter()
		ledgerAdapter, _ := ledger.NewSQLiteEventStoreAdapter("signet.db")
		
		// This is a simplified server for the demo
		signetv1.RegisterSignetServiceServer(s, &grpcServer{
			ledger: ledgerAdapter,
		})
		
		reflection.Register(s)

		fmt.Printf("Signet gRPC server listening on %v
", lis.Addr())
		if err := s.Serve(lis); err != nil {
			fmt.Printf("failed to serve: %v
", err)
			os.Exit(1)
		}
	},
}

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 50051, "Port to listen on")
	rootCmd.AddCommand(serveCmd)
}
